import React, { useContext, useEffect, useState } from 'react';
import {
  View,
  StyleSheet,
  ScrollView,
  TouchableOpacity,
  Text,
} from 'react-native';
import { ApplicationCard, IpInput } from '../components';
import { AppContext } from '../context/AppContext';
import { deleteToken, checkToken, getToken } from '../service';
import { globalStyles } from '../styles/global_style';

export default function SettingScreen() {
  const [token, setToken] = useState('');

  const {
    serverIp,
    aboutJson,
    isBlackTheme,
    setServerIp,
    setAboutJson,
    setIsConnected,
    setServicesConnected,
    setIsBlackTheme,
  } = useContext(AppContext);

  const checkIsToken = async () => {
    if ((await checkToken('token')) !== true) {
      setIsConnected(false);
    } else {
      await getToken('token', setToken);
    }
  };

  useEffect(() => {
    checkIsToken();
  }, []);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <ScrollView>
        <Text
          style={[isBlackTheme ? globalStyles.titleBlack : globalStyles.title, { textAlign: 'center', marginTop: 80 }]}
          accessibilityLabel="Setting">
          Setting
        </Text>
        <IpInput
          serverIp={serverIp}
          aboutJson={aboutJson}
          isBlackTheme={isBlackTheme}
          setServerIp={setServerIp}
          setAboutJson={setAboutJson}
          setServicesConnected={setServicesConnected}
        />
        <ApplicationCard
          isBlackTheme={isBlackTheme}
          setIsBlackTheme={setIsBlackTheme}
          setIsConnected={setIsConnected}
        />
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({});
