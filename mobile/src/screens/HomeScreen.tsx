import { useContext } from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { AppContext } from '../context/AppContext';
import { globalStyles } from '../styles/global_style';
import { IpInput } from '../components';
import { refreshServices } from '../service';

export default function HomeScreen() {
  const {
    isBlackTheme,
    serverIp,
    aboutJson,
    setAboutJson,
    setServerIp,
    setServicesConnected,
  } = useContext(AppContext);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}
          accessibilityLabel="Area51">
          Area51
        </Text>
        <View style={styles.textAlign}>
          <View style={styles.textAlign}>
            <Text
              style={
                isBlackTheme
                  ? globalStyles.subtitleBlack
                  : globalStyles.subtitle
              }
              accessibilityLabel="Automate">
              Automate
            </Text>
          </View>
          <View>
            <Text
              style={
                isBlackTheme
                  ? globalStyles.subtitleBlack
                  : globalStyles.subtitle
              }
              accessibilityLabel="Without Limits">
              without limits
            </Text>
          </View>
        </View>
        <View style={{ width: '90%' }}>
          <IpInput
            setAboutJson={setAboutJson}
            aboutJson={aboutJson}
            setServicesConnected={setServicesConnected}
            isBlackTheme={isBlackTheme}
            serverIp={serverIp}
            setServerIp={setServerIp}
          />
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  textAlign: {
    justifyContent: 'center',
    alignItems: 'center',
  },
});
