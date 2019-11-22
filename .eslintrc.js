module.exports = {
  env: {
    commonjs: true,
    es6: true,
    node: true
  },
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/eslint-recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking",
    "prettier",
    "prettier/@typescript-eslint",
    "plugin:vue/recommended",
    "@vue/prettier",
    "@vue/typescript",
    "@vue/typescript/recommended",
    "prettier/vue",
    "@vue/prettier/recommended"
  ],
  plugins: ["@typescript-eslint"],
  globals: {
    Atomics: "readonly",
    SharedArrayBuffer: "readonly"
  },
  rules: {
    "@typescript-eslint/explicit-function-return-type": "off"
  },
  overrides: [
    {
      files: ["workers/**/*.js"],
      rules: {
        "@typescript-eslint/no-var-requires": "off"
      }
    }
  ],
  parserOptions: {
    ecmaVersion: 2018,
    project: "tsconfig.json",
    parser: "@typescript-eslint/parser"
  }
};
