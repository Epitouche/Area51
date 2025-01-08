import React, { useContext, useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { Button } from 'react-native-paper';
import { WorkflowTable } from '../components';
import { AppContext } from '../context/AppContext';
import {
  parseServices,
  sendWorkflows,
  getAboutJson,
  deleteToken,
  checkToken,
  getToken,
} from '../service';
import { globalStyles } from '../styles/global_style';
import { Action, AppStackList, Reaction } from '../types';
import { NavigationProp, useNavigation } from '@react-navigation/native';

export default function WorkflowScreen() {
  const [token, setToken] = useState('');
  const [detailsModals, setdetailsModals] = useState(false);
  const [action, setAction] = useState<Action>({
    action_id: 0,
    name: '',
    description: '',
  });
  const [reaction, setReaction] = useState<Reaction>({
    reaction_id: 0,
    name: '',
    description: '',
  });

  const navigation = useNavigation<NavigationProp<AppStackList>>();

  const {
    serverIp,
    aboutJson,
    setAboutJson,
    setIsConnected,
    isBlackTheme,
    setServicesConnected,
    servicesConnected,
  } = useContext(AppContext);

  const handleLogout = () => {
    setIsConnected(false);
    deleteToken('token');
  };

  useEffect(() => {
    const checkIsToken = async () => {
      if ((await checkToken('token')) !== true) setIsConnected(false);
    };
    checkIsToken();
    if (aboutJson)
      parseServices({
        aboutJson,
        serverIp,
        setServicesConnected,
      });
  }, []);

  const handleSendWorkflow = async () => {
    await getToken('token', setToken);
    if (token !== 'Error: token not found' && action && reaction) {
      await sendWorkflows(token, serverIp, {
        action_id: action.action_id,
        reaction_id: reaction.reaction_id,
      });
      await getAboutJson(serverIp, setAboutJson);
      setAction({ action_id: 0, name: '', description: '' });
      setReaction({ reaction_id: 0, name: '', description: '' });
    }
  };

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <ScrollView>
        <View style={globalStyles.container}>
          <Text
            style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
            Dashboard
          </Text>
          <Button
            mode="contained"
            style={styles.loginButton}
            onPress={handleLogout}>
            <Text
              style={isBlackTheme ? globalStyles.textBlack : globalStyles.text}>
              Logout
            </Text>
          </Button>
          <View style={styles.buttonContainer}>
            <Button
              mode="contained"
              style={[styles.button, globalStyles.secondaryDark]}
              onPress={() => {
                navigation.navigate('ActionOrReaction', {
                  isAction: true,
                  setAction,
                });
              }}>
              {action.name === '' ? (
                <Text
                  style={
                    isBlackTheme ? globalStyles.textBlack : globalStyles.text
                  }>
                  Add Action
                </Text>
              ) : (
                <Text
                  style={
                    isBlackTheme ? globalStyles.textBlack : globalStyles.text
                  }>
                  {action.name}
                </Text>
              )}
            </Button>
            <Button
              mode="contained"
              style={[styles.button, globalStyles.secondaryDark]}
              onPress={() => {
                navigation.navigate('ActionOrReaction', {
                  isAction: false,
                  setReaction,
                });
              }}>
              {reaction.name === '' ? (
                <Text
                  style={
                    isBlackTheme ? globalStyles.textBlack : globalStyles.text
                  }>
                  Add Reaction
                </Text>
              ) : (
                <Text
                  style={
                    isBlackTheme ? globalStyles.textBlack : globalStyles.text
                  }>
                  {reaction.name}
                </Text>
              )}
            </Button>
          </View>
          <Button
            disabled={action.name === '' || reaction.name === ''}
            mode="contained"
            style={[styles.button, globalStyles.secondaryDark]}
            onPress={handleSendWorkflow}>
            <Text
              style={isBlackTheme ? globalStyles.textBlack : globalStyles.text}>
              Send Workflow
            </Text>
          </Button>
        </View>
        {aboutJson && (
          <View style={styles.tabContainer}>
            <WorkflowTable
              workflows={aboutJson.server.workflows}
              setDetailsModalVisible={setdetailsModals}
              detailsModalVisible={detailsModals}
            />
          </View>
        )}
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
  loginButton: {
    width: 'auto',
    backgroundColor: '#B454FD',
    justifyContent: 'center',
    alignItems: 'center',
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
  button: {
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
    marginTop: '10%',
    width: '100%',
    justifyContent: 'center',
  },
});
