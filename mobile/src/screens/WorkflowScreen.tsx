import React, { useContext, useEffect, useState } from 'react';
import {
  View,
  Text,
  StyleSheet,
  ScrollView,
  TouchableOpacity,
} from 'react-native';
import { WorkflowCard, WorkflowTab } from '../components';
import { AppContext } from '../context/AppContext';
import {
  parseServices,
  deleteToken,
  checkToken,
  getToken,
  getWorkflows,
} from '../service';
import { globalStyles } from '../styles/global_style';
import { AppStackList } from '../types';
import { NavigationProp, useNavigation } from '@react-navigation/native';

export default function WorkflowScreen() {
  const [token, setToken] = useState('');

  const navigation = useNavigation<NavigationProp<AppStackList>>();

  const {
    serverIp,
    aboutJson,
    setIsConnected,
    isBlackTheme,
    setServicesConnected,
    workflows,
    setWorkflows,
  } = useContext(AppContext);

  const handleLogout = () => {
    setIsConnected(false);
    deleteToken('token');
  };

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
          <TouchableOpacity
            style={[
              globalStyles.buttonFormat,
              isBlackTheme
                ? globalStyles.primaryLight
                : globalStyles.secondaryDark,
            ]}
            onPress={handleLogout}>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                globalStyles.textFormat,
              ]}
              accessibilityLabel="Logout">
              Logout
            </Text>
          </TouchableOpacity>
        </View>
        <WorkflowCard
          serverIp={serverIp}
          isBlackTheme={isBlackTheme}
          token={token}
          setWorkflows={setWorkflows}
        />
        <View style={styles.tabContainer}>
          <WorkflowTab
            workflows={workflows}
            isBlackTheme={isBlackTheme}
            navigation={navigation}
          />
        </View>
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  header: {
    fontSize: 32,
    color: '#222831',
    fontWeight: 'bold',
    marginTop: '20%',
  },
  textContainer: {
    justifyContent: 'center',
    alignItems: 'center',
  },
  buttonContainer: {
    marginTop: '10%',
    flexDirection: 'row',
    justifyContent: 'center',
    gap: '2%',
  },
  Actionbutton: {
    width: '48%',
    alignItems: 'center',
    justifyContent: 'center',
  },
  text: {
    color: '#222831',
    fontSize: 16,
    fontWeight: 'bold',
  },
  tabContainer: {
    width: '100%',
    marginTop: '1%',
    justifyContent: 'center',
    alignItems: 'center',
  },
});
