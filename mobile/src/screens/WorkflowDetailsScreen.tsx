import React, { useContext, useEffect, useState } from 'react';
import {
  StyleSheet,
  View,
  Text,
  TouchableOpacity,
  ScrollView,
} from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppStackList } from '../types';
import { AppContext } from '../context/AppContext';
import {
  NavigationProp,
  RouteProp,
  useNavigation,
  useRoute,
} from '@react-navigation/native';
import {
  deleteWorkflow,
  getReaction,
  getToken,
  getWorkflows,
  modifyWorkflows,
} from '../service';

type WorkflowDetailsProps = RouteProp<AppStackList, 'Workflow Details'>;

export default function WorkflowDetailsScreen() {
  const route = useRoute<WorkflowDetailsProps>();
  const { isBlackTheme, serverIp, setWorkflows } = useContext(AppContext);
  const { workflow } = route.params;
  const nav = useNavigation<NavigationProp<AppStackList>>();

  const [isToggled, setIsToggled] = useState(workflow.is_active);
  const [token, setToken] = useState('');
  const [reaction, setReaction] = useState<any>();

  const handleToggle = () => {
    setIsToggled(!isToggled);
  };

  useEffect(() => {
    const grabToken = async () => {
      await getToken('token', setToken);
    };
    grabToken();
  }, []);

  useEffect(() => {
    const grabReaction = async () => {
      if (token !== 'Error: token not found' && token !== '')
        await getReaction(serverIp, token, setReaction);
    };
    grabReaction();
  }, [token]);

  const handleSave = async () => {
    if (token !== 'Error: token not found' && token !== '') {
      await modifyWorkflows(serverIp, token, isToggled, workflow.workflow_id);
      await getWorkflows(serverIp, token, setWorkflows);
      nav.goBack();
    }
  };

  const handleDelete = async () => {
    if (token !== 'Error: token not found' && token !== '') {
      await deleteWorkflow(
        serverIp,
        token,
        workflow.workflow_id,
        workflow.name,
        workflow.action_id,
        workflow.reaction_id,
      );
      await getWorkflows(serverIp, token, setWorkflows);
      nav.goBack();
    }
  };

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <ScrollView>
        <View
          style={[
            styles.card,
            isBlackTheme
              ? globalStyles.secondaryLight
              : globalStyles.terciaryLight,
          ]}>
          <View style={{ justifyContent: 'center', alignItems: 'center' }}>
            <Text
              style={[
                isBlackTheme ? globalStyles.title : globalStyles.titleBlack,
              ]}
              accessibilityLabel={workflow.name}>
              {workflow.name.charAt(0).toUpperCase() + workflow.name.slice(1)}
            </Text>
          </View>

          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <Text
              style={[
                styles.bullet,
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
              ]}
              accessibilityLabel="Bullet">
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}
              accessibilityLabel="Action and Reaction">
              Action et Reaction
            </Text>
          </View>
          <View style={styles.buttonContainer}>
            <View
              style={[
                styles.ActionReactionButton,
                isBlackTheme
                  ? globalStyles.secondaryDark
                  : globalStyles.secondaryLight,
              ]}>
              <Text
                style={[
                  isBlackTheme
                    ? globalStyles.textColorBlack
                    : globalStyles.textColor,
                  globalStyles.textFormat,
                  styles.textFormat,
                ]}
                accessibilityLabel={workflow.action_name}
                numberOfLines={1}
                ellipsizeMode="tail">
                {workflow.action_name}
              </Text>
            </View>
            <View
              style={[
                styles.ActionReactionButton,
                isBlackTheme
                  ? globalStyles.secondaryDark
                  : globalStyles.secondaryLight,
              ]}>
              <Text
                style={[
                  isBlackTheme
                    ? globalStyles.textColorBlack
                    : globalStyles.textColor,
                  globalStyles.textFormat,
                  styles.textFormat,
                ]}
                accessibilityLabel={workflow.reaction_name}
                numberOfLines={1}
                ellipsizeMode="tail">
                {workflow.reaction_name}
              </Text>
            </View>
          </View>

          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <Text
              style={[
                styles.bullet,
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
              ]}
              accessibilityLabel="Bullet">
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}
              accessibilityLabel="Toggle Workflow">
              Active or disable the workflow
            </Text>
          </View>
          <View style={styles.toggleContainer}>
            <TouchableOpacity
              style={[
                styles.toggleButton,
                isToggled ? styles.toggledOn : styles.toggledOff,
              ]}
              onPress={handleToggle}>
              <Text
                style={[globalStyles.textFormat, globalStyles.textColorBlack]}
                accessibilityLabel={isToggled ? 'ON' : 'OFF'}>
                {isToggled ? 'ON' : 'OFF'}
              </Text>
            </TouchableOpacity>
          </View>
          <View
            style={{
              justifyContent: 'center',
              alignItems: 'center',
              flexDirection: 'row',
              gap: 20,
            }}>
            <TouchableOpacity
              style={[
                globalStyles.buttonFormat,
                { width: '48%', backgroundColor: 'red' },
              ]}
              onPress={handleDelete}>
              <Text
                style={[globalStyles.textFormat, globalStyles.textColorBlack]}
                accessibilityLabel={'Delete Workflows'}>
                Delete
              </Text>
            </TouchableOpacity>
            <TouchableOpacity
              style={[
                globalStyles.buttonFormat,
                isBlackTheme
                  ? globalStyles.secondaryDark
                  : globalStyles.secondaryLight,
                { width: '48%' },
              ]}
              onPress={handleSave}>
              <Text
                style={[
                  globalStyles.textFormat,
                  isBlackTheme
                    ? globalStyles.textColorBlack
                    : globalStyles.textColor,
                ]}
                accessibilityLabel={'Save Workflows modification'}>
                Save
              </Text>
            </TouchableOpacity>
          </View>
          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <Text
              style={[
                styles.bullet,
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
              ]}
              accessibilityLabel="Bullet">
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}
              accessibilityLabel="Toggle Workflow">
              Last reaction data
            </Text>
          </View>
          <View
            style={[
              styles.container,
              isBlackTheme
                ? globalStyles.secondaryLight
                : globalStyles.terciaryLight,
            ]}>
            {reaction &&
              reaction.map((item: any, index: number) => (
                <View key={index} style={styles.cardCode}>
                  <Text style={styles.bodyText}>{item.body}</Text>
                  <TouchableOpacity
                    onPress={() =>
                      console.log('URL clicked:', item.pull_request_url)
                    }>
                    <Text style={styles.urlText}>{item.pull_request_url}</Text>
                  </TouchableOpacity>
                </View>
              ))}
          </View>
        </View>
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    gap: 20,
    padding: 20,
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
    margin: 20,
    width: '90%',
  },
  subtitle: {
    fontSize: 16,
    fontWeight: '600',
  },
  input: {
    marginBottom: 20,
  },
  buttonContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  ActionReactionButton: {
    width: '48%',
    padding: 10,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
  },
  textFormat: {
    fontSize: 15,
  },
  disabledButton: {
    opacity: 0.5,
  },
  bullet: {
    fontSize: 20,
    marginRight: 10,
  },
  toggleContainer: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
  },
  toggleButton: {
    padding: 10,
    borderRadius: 20,
    alignItems: 'center',
    justifyContent: 'center',
    width: '100%',
  },
  toggledOn: {
    backgroundColor: '#4caf50',
  },
  toggledOff: {
    backgroundColor: '#f44336',
  },
  toggleText: {
    fontSize: 18,
  },
  container: {
    flex: 1,
    backgroundColor: '#f8f8f8',
    padding: 16,
  },
  cardCode: {
    backgroundColor: '#ffffff',
    padding: 16,
    borderRadius: 8,
    marginBottom: 16,
    shadowColor: '#000',
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  bodyText: {
    fontSize: 16,
    color: '#333',
    marginBottom: 8,
  },
  urlText: {
    fontSize: 14,
    color: '#1e90ff',
    textDecorationLine: 'underline',
  },
});
