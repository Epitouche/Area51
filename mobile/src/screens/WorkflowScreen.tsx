import React, { useContext, useEffect, useState } from 'react';
import { View, StyleSheet, ScrollView, Text } from 'react-native';
import { WorkflowCard, WorkflowTab } from '../components';
import { AppContext } from '../context/AppContext';
import {
  parseServices,
  checkToken,
  getToken,
  getWorkflows,
  refreshServices,
} from '../service';
import { globalStyles } from '../styles/global_style';
import { AppStackList } from '../types';
import { NavigationProp, useNavigation } from '@react-navigation/native';

export default function WorkflowScreen() {
  const [token, setToken] = useState('');
  const [refresh, setRefresh] = useState(false);

  const navigation = useNavigation<NavigationProp<AppStackList>>();

  const {
    serverIp,
    aboutJson,
    setIsConnected,
    isBlackTheme,
    setServicesConnected,
    workflows,
    setWorkflows,
    setAboutJson,
  } = useContext(AppContext);

  const grabWorkflows = async () => {
    if (token !== 'Error: token not found' && token !== '') {
      await getWorkflows(serverIp, token, setWorkflows);
    }
  };

  const checkIsToken = async () => {
    if ((await checkToken('token')) !== true) {
      setIsConnected(false);
    } else {
      await getToken('token', setToken);
    }
  };

  useEffect(() => {
    grabWorkflows();
    if (aboutJson)
      parseServices({
        aboutJson,
        serverIp,
        setServicesConnected,
      });
  }, [token]);

  useEffect(() => {
    if (refresh) {
      refreshServices({
        aboutJson,
        serverIp,
        setAboutJson,
        setServicesConnected,
      });
    }
  }, [refresh]);

  useEffect(() => {
    setTimeout(() => {
      checkIsToken();
    }, 300);
  }, []);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <ScrollView>
        <View style={globalStyles.container}>
          <Text
            style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}
            accessibilityLabel="Dashboard">
            Dashboard
          </Text>
          <WorkflowCard
            serverIp={serverIp}
            isBlackTheme={isBlackTheme}
            token={token}
            setWorkflows={setWorkflows}
            setRefresh={setRefresh}
          />
          <View style={styles.tabContainer}>
            <WorkflowTab
              workflows={workflows}
              isBlackTheme={isBlackTheme}
              navigation={navigation}
            />
          </View>
        </View>
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  tabContainer: {
    width: '100%',
    marginTop: '1%',
    justifyContent: 'center',
    alignItems: 'center',
  },
});
