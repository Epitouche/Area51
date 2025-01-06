import React, { useContext, useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { Button } from 'react-native-paper';
import { deleteToken, checkToken, getToken } from '../service/token';
import { ServicesModals } from '../components';
import { getAboutJson } from '../service';
import { AppContext } from '../context/AppContext';
import { sendWorkflows } from '../service/workflows';
import { globalStyles } from '../styles/global_style';
import { ActionReaction } from '../types';

export default function DashboardScreen() {
  const [token, setToken] = useState('');
  const [modalVisible, setModalVisible] = useState(false);
  const [isActionOrReaction, setIsActionOrReaction] = useState<boolean>(true);
  const [action, setAction] = useState<ActionReaction>({
    id: 0,
    name: '',
  });
  const [reaction, setReaction] = useState<ActionReaction>({
    id: 0,
    name: '',
  });

  const { serverIp, aboutjson, setAboutJson, setIsConnected, isBlackTheme } =
    useContext(AppContext);

  const handleLogout = () => {
    setIsConnected(false);
    deleteToken('token');
    deleteToken('github');
  };

  // const handleGithubLogin = async () => {
  //   if (github === 'Error: token not found') await githubLogin(serverIp);
  //   else {
  //     setServiceName('github');
  //     setDeconnectionModalVisible(!deconnectionModalVisible);
  //   }
  // };

  useEffect(() => {
    checkToken('token');
    checkToken('github');
    if (token === 'Error: token not found') {
      setIsConnected(false);
    }
  }, [token]);

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
                setIsActionOrReaction(true);
                setModalVisible(true);
              }}>
              <Text
                style={
                  isBlackTheme ? globalStyles.textBlack : globalStyles.text
                }>
                Add Action
              </Text>
            </Button>
            <Button
              mode="contained"
              style={[styles.button, globalStyles.buttonColor]}
              onPress={() => {
                setIsActionOrReaction(false);
                setModalVisible(true);
              }}>
              {reaction.id === 0 ? (
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
          <ServicesModals
            modalVisible={modalVisible}
            setModalVisible={setModalVisible}
            services={aboutjson}
            isAction={isActionOrReaction}
            setActionOrReaction={isActionOrReaction ? setAction : setReaction}
          />
        </View>
      </ScrollView>
    </View>
  );
}

{
  /* <View
  style={isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper}>
  <ScrollView>
    <View style={globalStyles.container}>
      <View style={styles.textContainer}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
          Dashboard
        </Text>
      </View>
      <Button
        mode="contained"
        style={styles.loginButton}
        onPress={handleLogout}>
        <Text style={styles.text}>Logout</Text>
      </Button>
      <OauthLoginButton
        handleOauthLogin={handleGithubLogin}
        color="#B454FD"
        name="Github"
        img="https://img.icons8.com/?size=100&id=12599&format=png"
      />
      <View style={styles.buttonContainer}>
        <Button
          mode="contained"
          style={styles.button}
          onPress={() => {
            setIsAction(true);
            setModalVisible(true);
          }}>
          <Text style={styles.text}>Add Action</Text>
        </Button>
        <Button
          mode="contained"
          style={styles.button}
          onPress={() => {
            setIsAction(false);
            setModalVisible(true);
          }}>
          <Text style={styles.text}>Add Reaction</Text>
        </Button>
      </View>
      <Button
        mode="contained"
        style={styles.button}
        onPress={handleSendWorkflow}>
        <Text style={styles.text}>Send Workflow</Text>
      </Button>
      {aboutjson && (
        <View style={styles.tabContainer}>
          <WorkflowTable
            workflows={aboutjson.server.workflows}
            setDetailsModalVisible={setdetailsModals}
            detailsModalVisible={detailsModals}
          />
        </View>
      )}
      {workflowsInfo && (
        <DetailsModals
          modalVisible={detailsModals}
          setModalVisible={setdetailsModals}
          workflows={workflowsInfo}
        />
      )}
      <ServicesModals
        modalVisible={modalVisible}
        setModalVisible={setModalVisible}
        services={aboutjson}
        isAction={isAction}
        setActionOrReaction={isAction ? setAction : setReaction}
      />
      <DeconnectionPopUp
        setModalVisible={setDeconnectionModalVisible}
        modalVisible={deconnectionModalVisible}
        service={serviceName}
      />
    </View>
  </ScrollView>
</View>; */
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
    width: '90%',
    justifyContent: 'center',
  },
});
