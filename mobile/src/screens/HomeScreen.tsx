import React, { useState, useContext, useEffect } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { AppContext } from '../context/AppContext';
import { getAboutJson } from '../service';
import { globalStyles } from '../styles/global_style';
import { IpInput } from '../components';

export default function HomeScreen() {
  const { serverIp, setAboutJson, isConnected, isBlackTheme } =
    useContext(AppContext);

  useEffect(() => {
    if (isConnected) getAboutJson(serverIp, setAboutJson);
  }, [isConnected, serverIp]);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
          Area51
        </Text>
        <View style={styles.textAlign}>
          <View style={styles.textAlign}>
            <Text
              style={
                isBlackTheme
                  ? globalStyles.subtitleBlack
                  : globalStyles.subtitle
              }>
              Automate
            </Text>
          </View>
          <View>
            <Text
              style={
                isBlackTheme
                  ? globalStyles.subtitleBlack
                  : globalStyles.subtitle
              }>
              without limits
            </Text>
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
