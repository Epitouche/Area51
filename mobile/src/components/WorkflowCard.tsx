import React, { useState } from 'react';
import {
  View,
  Text,
  TextInput,
  StyleSheet,
  TouchableOpacity,
} from 'react-native';
import { globalStyles } from '../styles/global_style';
import { Action, AppStackList, Reaction, Workflow } from '../types';
import { getWorkflows, sendWorkflows } from '../service';
import { NavigationProp, useNavigation } from '@react-navigation/native';

interface WorkflowCardProps {
  token: string;
  isBlackTheme?: boolean;
  serverIp: string;
  setWorkflows: (workflows: Workflow[]) => void;
}

export function WorkflowCard({
  isBlackTheme,
  serverIp,
  token,
  setWorkflows,
}: WorkflowCardProps) {
  const navigation = useNavigation<NavigationProp<AppStackList>>();

  const [workflowName, setWorkflowName] = useState('');
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

  const handleSendWorkflow = async () => {
    if (token !== 'Error: token not found' && action && reaction) {
      await sendWorkflows(token, serverIp, {
        action_id: action.action_id,
        reaction_id: reaction.reaction_id,
        name: workflowName,
      });
      await getWorkflows(serverIp, token, setWorkflows);
      setAction({ action_id: 0, name: '', description: '' });
      setReaction({ reaction_id: 0, name: '', description: '' });
      setWorkflowName('');
    }
  };

  const isDisabled = action.name === '' || reaction.name === '';

  return (
    <View
      style={[
        styles.card,
        isBlackTheme ? globalStyles.secondaryLight : globalStyles.secondaryDark,
      ]}>
      <Text
        style={[
          isBlackTheme ? globalStyles.title : globalStyles.titleBlack,
          styles.title,
        ]}>
        Create a Workflow
      </Text>
      <TextInput
        style={[
          isBlackTheme ? globalStyles.input : globalStyles.inputBlack,
          styles.input,
        ]}
        placeholder="Workflow Name"
        value={workflowName}
        onChangeText={setWorkflowName}
      />
      <Text
        style={[
          isBlackTheme ? globalStyles.title : globalStyles.titleBlack,
          styles.title,
          { marginBottom: 20 },
        ]}>
        Select a Action and a Reaction
      </Text>
      <View style={styles.buttonContainer}>
        <TouchableOpacity
          style={[
            styles.ActionReactionButton,
            isBlackTheme
              ? globalStyles.secondaryDark
              : globalStyles.secondaryLight,
          ]}
          onPress={() => {
            navigation.navigate('Options', { isAction: true, setAction });
          }}>
          {action.name === '' ? (
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColorBlack
                  : globalStyles.textColor,
                styles.textFormat,
              ]}
              numberOfLines={1}
              ellipsizeMode="tail">
              Add Action
            </Text>
          ) : (
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColorBlack
                  : globalStyles.textColor,
                styles.textFormat,
              ]}
              numberOfLines={1}
              ellipsizeMode="tail">
              {action.name}
            </Text>
          )}
        </TouchableOpacity>
        <TouchableOpacity
          style={[
            styles.ActionReactionButton,
            isBlackTheme
              ? globalStyles.secondaryDark
              : globalStyles.secondaryLight,
          ]}
          onPress={() => {
            navigation.navigate('Options', { isAction: false, setReaction });
          }}>
          {reaction.name === '' ? (
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColorBlack
                  : globalStyles.textColor,
                styles.textFormat,
              ]}
              numberOfLines={1}
              ellipsizeMode="tail">
              Add Reaction
            </Text>
          ) : (
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColorBlack
                  : globalStyles.textColor,
                styles.textFormat,
              ]}
              numberOfLines={1}
              ellipsizeMode="tail">
              {reaction.name}
            </Text>
          )}
        </TouchableOpacity>
      </View>
      <TouchableOpacity
        disabled={isDisabled}
        style={[
          globalStyles.buttonFormat,
          { marginTop: 20 },
          isBlackTheme
            ? globalStyles.secondaryDark
            : globalStyles.secondaryLight,
          isDisabled && styles.disabledButton,
        ]}
        onPress={handleSendWorkflow}>
        <Text
          style={[
            isBlackTheme ? globalStyles.textColorBlack : globalStyles.textColor,
            globalStyles.textFormat,
          ]}>
          Send Workflow
        </Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    padding: 20,
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
    margin: 20,
  },
  title: {
    fontSize: 20,
  },
  input: {
    marginBottom: 20,
  },
  buttonContainer: {
    flexDirection: 'row',
    gap: 10,
  },
  ActionReactionButton: {
    width: 140,
    padding: 10,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
    overflow: 'hidden',
  },
  textFormat: {
    fontSize: 15,
  },
  disabledButton: {
    opacity: 0.5,
  },
});