const { getDefaultConfig, mergeConfig } = require('@react-native/metro-config');

/**
 * Metro configuration
 * https://reactnative.dev/docs/metro
 *
 * @type {import('metro-config').MetroConfig}
 */
const config = {
  transformer: {
    babelTransformerPath: require.resolve('react-native-svg-transformer'),
  },
  resolver: {
    assetExts: [
      'bin',
      'txt',
      'png',
      'jpg',
      'gif',
      'ttf',
      'otf',
      'ttc',
      'mp4',
      'mp3',
      'wav',
      'svg',
    ],
    sourceExts: ['js', 'jsx', 'ts', 'tsx', 'json', 'svg'],
  },
};

module.exports = mergeConfig(getDefaultConfig(__dirname), config);
