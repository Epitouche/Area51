import React, { useState, useContext, useEffect } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { AppContext } from '../context/AppContext';
import { getAboutJson } from '../service';
import { globalStyles } from '../styles/global_style';
import { IpInput } from '../components';

export default function HomeScreen() {
  const { serverIp, setAboutJson, isConnected } = useContext(AppContext);
  const [ipTmp, setIpTmp] = useState('');

  useEffect(() => {
    if (isConnected) getAboutJson(serverIp, setAboutJson);
  }, [isConnected, serverIp]);

  return (
    <View style={globalStyles.wallpaper}>
      <View style={globalStyles.container}>
        <Text style={globalStyles.titleWhite}>Area51</Text>
        <View style={styles.textAlign}>
          <View style={styles.textAlign}>
            <Text style={globalStyles.subtitleWhite}>Automate</Text>
          </View>
          <View>
            <Text style={globalStyles.subtitleWhite}>without limits</Text>
          </View>
        </View>
        <IpInput />
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
