import { Text, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppContext } from '../context/AppContext';
import { useContext } from 'react';

export default function ServiceScreen() {
  const { isBlackTheme } = useContext(AppContext);
  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
          Service Screen
        </Text>
      </View>
    </View>
  );
}
