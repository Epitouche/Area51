import React, { useContext, useEffect, useState } from 'react';
import { View, StyleSheet, ScrollView } from 'react-native';
import { IpInput } from '../components';
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
  } = useContext(AppContext);

  const handleLogout = () => {
    setIsConnected(false);
    deleteToken('token');
  };

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
        <View style={globalStyles.container}>
          <IpInput
            serverIp={serverIp}
            aboutJson={aboutJson}
            isBlackTheme={isBlackTheme}
            setServerIp={setServerIp}
            setAboutJson={setAboutJson}
            setServicesConnected={setServicesConnected}
          />
        </View>
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
});
