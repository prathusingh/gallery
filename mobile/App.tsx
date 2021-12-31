import React from 'react';
import { Image, StatusBar, StyleSheet, View } from 'react-native';
import { ImageManipulator, Asset, AppLoading } from 'expo';

import BackgroundImage from './components/BackgroundImage';

interface Size {
  width: number;
  height: number;
}

const getSize = (uri: string): Promise<Size> => new Promise(
  (resolve, reject) => Image.getSize(uri, (width, height) => resolve({width, height}), reject)
)

const screens = [
  require("./assets/stories/story1.png"),
  require("./assets/stories/story2.png"),
  require("./assets/stories/story3.png"),
];


interface IAppProps {}
interface IAppState {
  stories: { top: string, bottom: string }[];
  index: number;
}

export default class App extends React.Component<IAppProps, IAppState {
  state: IAppState = {
    stories: [],
    index: 1
  };

  async componentDidMount() {
    const edits = screens.map(async (screen) => {
      const image = Asset.fromModule(screen);
      await image.downloadAsync();
      const { localUri } = image;
      const { width, height } = await getSize(localUri);

      const cropTop = {
        crop: {
          originX: 0,
          originY: 0,
          width,
          height: height / 2,
        },
      };

      const cropBottom = {
        crop: {
          originX: 0,
          originY: height / 2,
          width,
          height: height / 2
        },
      };

      const { uri: top } = await ImageManipulator.manipulateAsync(localUri, [cropTop]);
      const { uri: bottom } = await ImageManipulator.manipulateAsync(localUri, [cropBottom]);
      return { top, bottom };
    });

    const stories = await Promise.all(edits);
    this.setState({ stories });
  }

  render() {
    const { stories, index } = this.state;
    if (stories.length === 0) {
      return (
        <AppLoading />
      )
    }

    const prev = stories[index - 1];
    const current = stories[index];
    const next = stories[index + 1];

    return (
      <View style={styles.container}>
        <StatusBar hidden />
        <BackgroundImage top={prev.top} bottom={next.bottom} />
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});
