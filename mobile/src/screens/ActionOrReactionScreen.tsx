import React, { useContext, useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { AppContext } from '../context/AppContext';
import { globalStyles } from '../styles/global_style';
import { Action, AppStackList, Reaction, ServicesParse } from '../types';
import { RouteProp, useNavigation, useRoute } from '@react-navigation/native';
import { Button } from 'react-native-paper';

type ActionOrReactionProps = RouteProp<AppStackList, 'Options'>;

function NoService() {
  const { isBlackTheme } = useContext(AppContext);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
          No Service Connected
        </Text>
      </View>
    </View>
  );
}

function ActionOrReaction() {
  const navigation = useNavigation();
  const defaultService = {
    name: '',
    isConnected: false,
    actions: [],
    reactions: [],
  };

  const route = useRoute<ActionOrReactionProps>();
  const [selectedService, setSelectedService] =
    useState<ServicesParse>(defaultService);
  const [selectedActionOrReactionId, setSelectedActionOrReactionId] = useState<
    Action | Reaction
  >();
  const { isAction, setAction, setReaction } = route.params;
  const { servicesConnected, isBlackTheme } = useContext(AppContext);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={styles.flexContainer}>
        <View style={globalStyles.container}>
          <Text
            style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
            {isAction ? 'Creating an Action' : 'Creating an Reaction'}
          </Text>
          <Text
            style={[
              isBlackTheme
                ? globalStyles.textColorBlack
                : globalStyles.textColor,
              globalStyles.textFormat,
            ]}>
            Select a service
          </Text>
          {servicesConnected.services.map((service, index) => {
            if (service.isConnected) {
              return (
                <Button
                  key={index}
                  mode="contained"
                  style={
                    isBlackTheme
                      ? globalStyles.primaryLight
                      : globalStyles.secondaryDark
                  }
                  onPress={() => {
                    setSelectedService(service);
                  }}>
                  <Text
                    style={[
                      isBlackTheme
                        ? globalStyles.textColor
                        : globalStyles.textColorBlack,
                      globalStyles.textFormat,
                    ]}>
                    {service.name}
                  </Text>
                </Button>
              );
            }
            return null;
          })}
          {selectedService && (
            <View style={styles.textContainer}>
              <Text
                style={[
                  isBlackTheme
                    ? globalStyles.textColorBlack
                    : globalStyles.textColor,
                  globalStyles.textFormat,
                ]}>
                {isAction ? 'Select an Action for ' : 'Select an Reaction for '}
                {selectedService.name}
              </Text>
              {isAction
                ? selectedService.actions.map((action, index) => {
                    if (setAction) {
                      return (
                        <Button
                          key={index}
                          mode="contained"
                          style={
                            isBlackTheme
                              ? globalStyles.primaryLight
                              : globalStyles.secondaryDark
                          }
                          onPress={() => {
                            setSelectedActionOrReactionId(action);
                          }}>
                          <Text
                            style={[
                              isBlackTheme
                                ? globalStyles.textColor
                                : globalStyles.textColorBlack,
                              globalStyles.textFormat,
                            ]}>
                            {action.name}
                          </Text>
                        </Button>
                      );
                    }
                  })
                : selectedService.reactions.map((reaction, index) => {
                    if (setReaction) {
                      return (
                        <Button
                          key={index}
                          mode="contained"
                          style={
                            isBlackTheme
                              ? globalStyles.primaryLight
                              : globalStyles.secondaryDark
                          }
                          onPress={() => {
                            setSelectedActionOrReactionId(reaction);
                          }}>
                          <Text
                            style={[
                              isBlackTheme
                                ? globalStyles.textColor
                                : globalStyles.textColorBlack,
                              globalStyles.textFormat,
                            ]}>
                            {reaction.name}
                          </Text>
                        </Button>
                      );
                    }
                  })}
            </View>
          )}
        </View>
        <View style={styles.containerSaveButton}>
          <Button
            style={[
              styles.saveButton,
              isBlackTheme
                ? globalStyles.primaryLight
                : globalStyles.secondaryDark,
            ]}
            onPress={() => {
              if (isAction) {
                if (selectedActionOrReactionId)
                  setAction && setAction(selectedActionOrReactionId as Action);
              } else {
                if (selectedActionOrReactionId)
                  setReaction &&
                    setReaction(selectedActionOrReactionId as Reaction);
              }
              navigation.goBack();
            }}>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                globalStyles.textFormat
              ]}>
              Save
            </Text>
          </Button>
        </View>
      </View>
    </View>
  );
}

export default function ActionOrReactionScreen() {
  const { servicesConnected } = useContext(AppContext);
  const [connected, setConneted] = useState(0);
  useEffect(() => {
    if (servicesConnected.services)
      servicesConnected.services.map(service => {
        if (service.isConnected) setConneted(connected + 1);
      });
  }, []);

  return connected > 0 ? <ActionOrReaction /> : <NoService />;
}

const styles = StyleSheet.create({
  flexContainer: {
    flex: 1,
    justifyContent: 'space-between',
    width: '90%',
  },
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
  button: {
    width: 'auto',
    backgroundColor: '#F7FAFB',
    justifyContent: 'center',
    alignItems: 'center',
  },
  buttonSelect: {
    width: 'auto',
    backgroundColor: 'red',
    justifyContent: 'center',
    alignItems: 'center',
  },
  textContainer: {
    gap: 20,
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
  containerSaveButton: {
    alignItems: 'center',
    justifyContent: 'center',
  },
  saveButton: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 20,
    marginBottom: 40,
  },
});
