import React, { useState, useContext, useEffect } from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { AppContext } from '../context/AppContext';
import { globalStyles } from '../styles/global_style';
import { IpInput } from '../components';
import { getAboutJson, parseServices } from '../service';

export default function HomeScreen() {
  const {
    isBlackTheme,
    serverIp,
    aboutJson,
    setAboutJson,
    setServicesConnected,
  } = useContext(AppContext);

  const handleRefresh = async () => {
    if (serverIp != '') {
      await getAboutJson(serverIp, setAboutJson);
      if (aboutJson)
        await parseServices({ serverIp, aboutJson, setServicesConnected });
    }
  };

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
        <IpInput />
        <TouchableOpacity
          style={[
            globalStyles.buttonFormat,
            isBlackTheme ? globalStyles.primaryLight : globalStyles.primaryDark,
          ]}
          onPress={handleRefresh}>
          <Text
            style={[
              isBlackTheme
                ? globalStyles.textColor
                : globalStyles.textColorBlack,
              globalStyles.textFormat,
            ]}>
            Refresh
          </Text>
        </TouchableOpacity>
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
