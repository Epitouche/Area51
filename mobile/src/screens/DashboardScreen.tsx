import React, { useContext, useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { Button } from 'react-native-paper';
import { deleteToken, checkToken, getToken } from '../service/token';
import { ServicesModals, WorkflowTable } from '../components';
import { AppContext } from '../context/AppContext';
import { parseServices, sendWorkflows, getAboutJson } from '../service';
import { globalStyles } from '../styles/global_style';
import { ActionReaction, AppStackList } from '../types';
import { NavigationProp, useNavigation } from '@react-navigation/native';

export default function DashboardScreen() {
  const [token, setToken] = useState('');
  const [modalVisible, setModalVisible] = useState(false);
  const [isActionOrReaction, setIsActionOrReaction] = useState<boolean>(true);
  const [detailsModals, setdetailsModals] = useState(false);
  const [action, setAction] = useState<ActionReaction>({
    id: 0,
    name: '',
  });
  const [reaction, setReaction] = useState<ActionReaction>({
    id: 0,
    name: '',
  });

  const navigation = useNavigation<NavigationProp<AppStackList>>();

  const {
    serverIp,
    aboutJson,
    setAboutJson,
    setIsConnected,
    isBlackTheme,
    setServicesConnected,
    servicesConnected
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

  console.log('servicesConnected', servicesConnected);

  const handleSendWorkflow = async () => {
    await getToken('token', setToken);
    if (token !== 'Error: token not found' && action && reaction) {
      await sendWorkflows(token, serverIp, {
        action_id: action.id,
        reaction_id: reaction.id,
      });
      await getAboutJson(serverIp, setAboutJson);
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
              style={[styles.button, globalStyles.buttonColor]}
              onPress={() => {
                navigation.navigate('ActionOrReaction', {
                  isAction: true
                });
                // setIsActionOrReaction(true);
                // setModalVisible(true);
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
              style={[styles.button, globalStyles.buttonColor]}
              onPress={() => {
                setIsActionOrReaction(false);
                setModalVisible(true);
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
          {aboutJson && (
            <Button
              //disabled={action.name === '' || reaction.name === ''}
              mode="contained"
              style={[styles.button, globalStyles.buttonColor]}
              onPress={() =>
                parseServices({
                  aboutJson,
                  serverIp,
                  setServicesConnected,
                })
              }>
              <Text
                style={
                  isBlackTheme ? globalStyles.textBlack : globalStyles.text
                }>
                Send Workflow
              </Text>
            </Button>
          )}
          <ServicesModals
            modalVisible={modalVisible}
            setModalVisible={setModalVisible}
            services={aboutJson}
            isAction={isActionOrReaction}
            setActionOrReaction={isActionOrReaction ? setAction : setReaction}
          />
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
  container: {
    width: '90%',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 20,
  },
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
    marginTop: 20,
    flexDirection: 'row',
    justifyContent: 'center',
    gap: 20,
    width: '90%',
  },
  button: {
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
    justifyContent: 'center',
  },
});
