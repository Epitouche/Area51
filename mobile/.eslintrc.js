module.exports = {
  root: true,
  extends: ['@react-native', 'plugin:prettier/recommended'],
  rules: {
    'prettier/prettier': [
      'error',
      {
        arrowParens: 'avoid',
        bracketSameLine: true,
        bracketSpacing: true,
        singleQuote: true,
        trailingComma: 'all',
      },
    ],
    'padding-line-between-statements': [
      'error',
      { blankLine: 'always', prev: '*', next: 'export' },
    ],
    curly: ['error', 'multi-or-nest'],
  },
};
