module.exports = {
  env: {
    commonjs: true,
    es6: true,
    node: true
  },
  extends: ["standard", "prettier", "plugin:prettier/recommended", "plugin:vue/essential",
  "@vue/prettier", "@vue/typescript"],
  globals: {
    Atomics: "readonly",
    SharedArrayBuffer: "readonly"
  },
  parserOptions: {
    ecmaVersion: 2018,
    parser: "@typescript-eslint/parser"
  }
};
