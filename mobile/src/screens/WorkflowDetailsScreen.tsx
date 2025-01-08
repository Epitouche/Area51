import React, { useContext, useEffect, useState } from 'react';
import {
  StyleSheet,
  View,
  Text,
  TouchableOpacity,
  ScrollView,
} from 'react-native';
import {} from 'react-native-paper';
import { globalStyles } from '../styles/global_style';
import { AppStackList } from '../types';
import { AppContext } from '../context/AppContext';
import { RouteProp, useRoute } from '@react-navigation/native';

type WorkflowDetailsProps = RouteProp<AppStackList, 'Workflow Details'>;

export default function WorkflowDetailsScreen() {
  const { isBlackTheme } = useContext(AppContext);
  const route = useRoute<WorkflowDetailsProps>();
  const { workflow } = route.params;

  const [isToggled, setIsToggled] = useState(workflow.is_active);

  const handleToggle = () => {
    setIsToggled(!isToggled);
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
              : globalStyles.secondaryDark,
          ]}>
          <View style={{ justifyContent: 'center', alignItems: 'center' }}>
            <Text
              style={[
                isBlackTheme ? globalStyles.title : globalStyles.titleBlack,
              ]}>
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
              ]}>
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}>
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
                ]}>
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
                ]}>
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
              ]}>
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}>
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
                style={[
                  globalStyles.textFormat,
                  isToggled
                    ? globalStyles.textColorBlack
                    : globalStyles.textColor,
                ]}>
                {isToggled ? 'ON' : 'OFF'}
              </Text>
            </TouchableOpacity>
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
    backgroundColor: 'green',
  },
  toggledOff: {
    backgroundColor: 'red',
  },
  toggleText: {
    fontSize: 18,
  },
});
