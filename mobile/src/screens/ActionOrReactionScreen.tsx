import React, { useContext, useEffect, useState } from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { AppContext } from '../context/AppContext';
import { globalStyles } from '../styles/global_style';
import { AppStackList, ServicesParse } from '../types';
import { RouteProp, useRoute } from '@react-navigation/native';
import { Button } from 'react-native-paper';

type LoginScreenRouteProp = RouteProp<AppStackList, 'ActionOrReaction'>;

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
  const route = useRoute<LoginScreenRouteProp>();
  const [selectedService, setSelectedService] = useState<ServicesParse>();
  const [selectedActionOrReactionId, setSelectedActionOrReactionId] = useState<number>();
  const { isAction } = route.params;
  const { servicesConnected, isBlackTheme } = useContext(AppContext);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <ScrollView>
        <View style={globalStyles.container}>
          <Text
            style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
            {isAction ? 'Creating an Action' : 'Creating an Reaction'}
          </Text>
          <Text
            style={isBlackTheme ? globalStyles.textBlack : globalStyles.text}>
            Select a service
          </Text>
          {servicesConnected.services.map((service, index) => {
            if (selectedService)
              console.log(selectedService.name, service.name)
            if (service.isConnected) {
              return (
                <Button
                  key={index}
                  mode="contained"
                  style={
                    selectedService && selectedService.name === service.name
                      ? styles.buttonSelect
                      : styles.button
                  }
                  onPress={() => {
                    setSelectedService(service);
                  }}>
                  <Text
                    style={
                      isBlackTheme ? globalStyles.textBlack : globalStyles.text
                    }>
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
                style={
                  isBlackTheme ? globalStyles.textBlack : globalStyles.text
                }>
                {isAction ? 'Select an Action for ' : 'Select an Reaction for '}
                {selectedService.name}
              </Text>
              {isAction
                ? selectedService.actions.map((action, index) => (
                    <Button
                      key={index}
                      mode="contained"
                      style={
                        selectedActionOrReactionId === action.action_id
                          ? styles.buttonSelect
                          : styles.button
                      }
                      onPress={() => {
                        console.log('Create action/reaction');
                      }}>
                      <Text
                        style={
                          isBlackTheme
                            ? globalStyles.textBlack
                            : globalStyles.text
                        }>
                        {action.name}
                      </Text>
                    </Button>
                  ))
                : selectedService.reactions.map((reaction, index) => (
                    <Button
                      key={index}
                      mode="contained"
                      style={styles.button}
                      onPress={() => {
                        console.log('Create action/reaction');
                      }}>
                      <Text
                        style={
                          isBlackTheme
                            ? globalStyles.textBlack
                            : globalStyles.text
                        }>
                        {reaction.name}
                      </Text>
                    </Button>
                  ))}
            </View>
          )}
        </View>
      </ScrollView>
    </View>
  );
}

export default function ActionOrReactionScreen() {
  const { servicesConnected, isBlackTheme } = useContext(AppContext);
  const [connected, setConneted] = useState(0);
  useEffect(() => {
    if (servicesConnected.services.length > 0)
      servicesConnected.services.map(service => {
        if (service.isConnected) setConneted(connected + 1);
      });
  }, []);
  console.log(servicesConnected.services.length);

  return connected > 0 ? <ActionOrReaction /> : <NoService />;
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
