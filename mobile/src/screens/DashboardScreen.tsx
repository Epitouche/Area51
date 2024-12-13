import React, { useContext, useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { Button } from 'react-native-paper';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import { deleteToken, checkToken, getToken } from '../service/token';
import {
  DeconnectionPopUp,
  DetailsModals,
  GithubLoginButton,
  ServicesModals,
  WorkflowTable,
} from '../components';
import { getAboutJson, githubLogin } from '../service';
import { AppContext } from '../context/AppContext';
import { getWorkflows, sendWorkflows } from '../service/workflows';
import { PullRequestComment } from '../types';

type DashboardNavigationProp = StackNavigationProp<
  RootStackParamList,
  'Dashboard'
>;

type Props = {
  navigation: DashboardNavigationProp;
};

export default function DashboardScreen({ navigation }: Props) {
  const [token, setToken] = useState('');
  const [github, setGithub] = useState('');
  const [modalVisible, setModalVisible] = useState(false);
  const [deconnectionModalVisible, setDeconnectionModalVisible] =
    useState(false);
  const [serviceName, setServiceName] = useState('');
  const [isAction, setIsAction] = useState(true);
  const [action, setAction] = useState<number>(1);
  const [reaction, setReaction] = useState<number>(1);
  const [detailsModals, setdetailsModals] = useState(false);
  const [workflowsInfo, setWorkflowsInfo] = useState<PullRequestComment[]>();

  const { serverIp, aboutjson, setAboutJson } = useContext(AppContext);

  const handleLogout = () => {
    deleteToken('token');
    deleteToken('github');
    navigation.navigate('Home');
  };

  const handleGithubLogin = async () => {
    if (github === 'Error: token not found')
      await githubLogin(serverIp, setGithub);
    else {
      setServiceName('github');
      setDeconnectionModalVisible(!deconnectionModalVisible);
    }
  };

  useEffect(() => {
    if (token !== 'Error: token not found')
      getWorkflows(serverIp, token, setWorkflowsInfo);
  }, [detailsModals]);

  useEffect(() => {
    checkToken('token');
    checkToken('github');
    if (
      token === 'Error: token not found' &&
      github === 'Error: token not found'
    ) {
      navigation.navigate('Home');
    }
  }, [token]);

  const handleSendWorkflow = async () => {
    await getToken('token', setToken);
    if (token !== 'Error: token not found' && action && reaction) {
      await sendWorkflows(token, serverIp, {
        action_id: action,
        reaction_id: reaction,
      });
      await getAboutJson(serverIp, setAboutJson);
    }
  };

  return (
    <ScrollView>
      <View style={styles.container}>
        <View style={styles.textContainer}>
          <Text style={styles.header}>Dashboard</Text>
        </View>
        <Button
          mode="contained"
          style={styles.loginButton}
          onPress={handleLogout}>
          <Text style={styles.text}>Logout</Text>
        </Button>
        <GithubLoginButton
          handleGithubLogin={handleGithubLogin}
          color="#B454FD"
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
    width: '48%',
    backgroundColor: '#B454FD',
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
