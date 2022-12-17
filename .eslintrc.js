module.exports = {
  root: true,
  parser: "@typescript-eslint/parser",
  plugins: ["@typescript-eslint"],
  settings: {
    // This is needed to make eslint happy with name aliases
    "import/resolver": {
      "babel-module": {},
    },
  },
  overrides: [
    {
      files: ["*.ts", "*.tsx"],
      rules: {
        "@typescript-eslint/no-shadow": ["error"],
        "no-shadow": "off",
        "no-undef": "off",
      },
    },
  ],
}
