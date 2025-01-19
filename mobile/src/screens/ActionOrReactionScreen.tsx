import React, { useContext, useEffect, useState } from 'react';
import {
  View,
  Text,
  StyleSheet,
  ScrollView,
  TouchableOpacity,
  TextInput,
} from 'react-native';
import { AppContext } from '../context/AppContext';
import { globalStyles } from '../styles/global_style';
import { AppStackList, ServicesParse } from '../types';
import { RouteProp, useNavigation, useRoute } from '@react-navigation/native';

interface Values {
  id: number;
  name: string;
  description: string;
  options: { [key: string]: any };
}

type Options = {
  [key: string]: any;
};

type ActionOrReactionProps = RouteProp<AppStackList, 'Options'>;

interface OptionsInputProps {
  options: Options;
  onChange: (key?: string, value?: string) => void;
  isBlackTheme?: boolean;
}

function OptionsInput({ options, onChange, isBlackTheme }: OptionsInputProps) {
  const renderOptions = (
    options: { [key: string]: any },
    parentKey?: string,
  ) => {
    if (typeof options === 'string') {
      return (
        <View>
          <Text
            style={[
              isBlackTheme
                ? globalStyles.textColor
                : globalStyles.textColorBlack,
              globalStyles.textFormat,
            ]}>
            {parentKey
              ? (parentKey.split('.').pop()?.charAt(0).toUpperCase() ?? '') +
                (parentKey.split('.').pop()?.slice(1) ?? '')
              : ''}
          </Text>
          <TextInput
            style={isBlackTheme ? globalStyles.input : globalStyles.inputBlack}
            value={options}
            onChangeText={text => onChange(parentKey, text)}
            autoCapitalize="none"
            placeholder={options || `Enter ${parentKey}`}
            placeholderTextColor={isBlackTheme ? '#0a0a0a' : 'f5f5f5'}
          />
        </View>
      );
    } else if (typeof options === 'object' && options !== null) {
      return Object.keys(options).map(key => (
        <View key={key}>
          {renderOptions(options[key], parentKey ? `${parentKey}.${key}` : key)}
        </View>
      ));
    }
    return null;
  };

  return <View>{renderOptions(options)}</View>;
}

function NoService() {
  const { isBlackTheme } = useContext(AppContext);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}
          accessibilityLabel="No Service">
          No Service Connected
        </Text>
      </View>
    </View>
  );
}

function ActionOrReaction() {
  const navigation = useNavigation();
  const defaultService: ServicesParse = {
    name: '',
    isConnected: false,
    actions: [],
    reactions: [],
    image: '',
    description: '',
    is_oauth: false,
  };

  const route = useRoute<ActionOrReactionProps>();
  const [selectedService, setSelectedService] =
    useState<ServicesParse>(defaultService);
  const [selectedActionOrReactionId, setSelectedActionOrReactionId] =
    useState<Values>();
  const [options, setOptions] = useState<{ [key: string]: any }>({});
  const { isAction, setValues } = route.params;
  const { servicesConnected, isBlackTheme } = useContext(AppContext);

  useEffect(() => {
    setSelectedActionOrReactionId(undefined);
    setOptions({});
  }, [selectedService]);

  useEffect(() => {
    if (selectedActionOrReactionId) {
      setOptions(selectedActionOrReactionId.options || {});
    }
  }, [selectedActionOrReactionId]);

  const handleOptionsChange = (key: string | undefined, value?: string) => {
    let updatedOptions = { ...options };

    const updateNestedValue = (obj: any, key: string, value: string): any => {
      const keys = key.split('.');
      const lastKey = keys.pop();

      if (keys.length > 0) {
        const parentObj = keys.reduce(
          (acc, currentKey) => acc[currentKey],
          obj,
        );
        if (parentObj && typeof parentObj === 'object') {
          parentObj[lastKey as string] = value;
        }
      } else if (lastKey) {
        obj[lastKey] = value;
      }

      return obj;
    };

    if (key) {
      if (typeof updatedOptions === 'object') {
        updatedOptions = updateNestedValue(updatedOptions, key, value || '');
      }
    } else {
      updatedOptions = { ...updatedOptions, value: value || '' };
    }

    setOptions(updatedOptions);
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
          <Text
            style={isBlackTheme ? globalStyles.title : globalStyles.titleBlack}
            accessibilityLabel={
              isAction ? 'Creating Action' : 'Creating Reaction'
            }>
            {isAction ? 'Creating an Action' : 'Creating an Reaction'}
          </Text>
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
              accessibilityLabel="Select Service">
              Select a service
            </Text>
          </View>
          <View style={styles.buttonContainer}>
            {servicesConnected.services.map((service, index) => {
              if (service.isConnected) {
                return (
                  <TouchableOpacity
                    key={index}
                    style={[
                      globalStyles.buttonFormat,
                      isBlackTheme
                        ? globalStyles.primaryDark
                        : globalStyles.secondaryLight,
                    ]}
                    onPress={() => {
                      setSelectedService(service);
                    }}>
                    <Text
                      style={[
                        isBlackTheme
                          ? globalStyles.textColorBlack
                          : globalStyles.textColor,
                        globalStyles.textFormat,
                      ]}
                      accessibilityLabel={service.name}>
                      {service.name}
                    </Text>
                  </TouchableOpacity>
                );
              }
              return null;
            })}
          </View>
          {selectedService && selectedService.name !== '' && (
            <View style={styles.textContainer}>
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
                  accessibilityLabel={
                    isAction ? 'Select Action' : 'Select Reaction'
                  }>
                  {isAction
                    ? 'Select an Action for '
                    : 'Select an Reaction for '}
                  {selectedService.name}
                </Text>
              </View>
              {isAction ? (
                selectedService.actions &&
                selectedService.actions.length > 0 ? (
                  selectedService.actions.map((action, index) => {
                    return (
                      <TouchableOpacity
                        key={index}
                        style={[
                          globalStyles.buttonFormat,
                          isBlackTheme
                            ? globalStyles.secondaryDark
                            : globalStyles.primaryLight,
                        ]}
                        onPress={() => {
                          setSelectedActionOrReactionId({
                            id: action.action_id,
                            name: action.name,
                            description: action.description,
                            options: action.options,
                          });
                        }}>
                        <Text
                          style={[
                            isBlackTheme
                              ? globalStyles.textColorBlack
                              : globalStyles.textColor,
                            globalStyles.textFormat,
                          ]}
                          accessibilityLabel={action.name}>
                          {action.name}
                        </Text>
                      </TouchableOpacity>
                    );
                  })
                ) : (
                  <Text
                    style={[
                      globalStyles.textColorBlack,
                      globalStyles.textFormat,
                      { justifyContent: 'center' },
                    ]}
                    accessibilityLabel={'No Action Available'}>
                    No Action Available
                  </Text>
                )
              ) : selectedService.reactions &&
                selectedService.reactions.length > 0 ? (
                selectedService.reactions.map((reaction, index) => {
                  return (
                    <TouchableOpacity
                      key={index}
                      style={[
                        globalStyles.buttonFormat,
                        isBlackTheme
                          ? globalStyles.secondaryDark
                          : globalStyles.primaryLight,
                      ]}
                      onPress={() => {
                        setSelectedActionOrReactionId({
                          id: reaction.reaction_id,
                          name: reaction.name,
                          description: reaction.description,
                          options: reaction.options,
                        });
                      }}>
                      <Text
                        style={[
                          isBlackTheme
                            ? globalStyles.textColorBlack
                            : globalStyles.textColor,
                          globalStyles.textFormat,
                        ]}
                        accessibilityLabel={reaction.name}>
                        {reaction.name}
                      </Text>
                    </TouchableOpacity>
                  );
                })
              ) : (
                <Text
                  style={[
                    globalStyles.textColorBlack,
                    globalStyles.textFormat,
                    { justifyContent: 'center' },
                  ]}
                  accessibilityLabel={'No Reaction Available'}>
                  No Reaction Available
                </Text>
              )}
              {selectedActionOrReactionId?.options && (
                <>
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
                      accessibilityLabel={'Enter the Options'}>
                      Enter the Options
                    </Text>
                  </View>
                  <OptionsInput
                    options={options}
                    onChange={handleOptionsChange}
                    isBlackTheme={isBlackTheme}
                  />
                </>
              )}
            </View>
          )}
          <TouchableOpacity
            style={[
              globalStyles.buttonFormat,
              isBlackTheme
                ? globalStyles.secondaryDark
                : globalStyles.primaryLight,
            ]}
            onPress={() => {
              if (selectedActionOrReactionId) {
                setValues({
                  id: selectedActionOrReactionId.id,
                  name: selectedActionOrReactionId.name,
                  description: selectedActionOrReactionId.description,
                  options: options,
                });
              }
              navigation.goBack();
            }}>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColorBlack
                  : globalStyles.textColor,
                globalStyles.textFormat,
              ]}
              accessibilityLabel="Save">
              Save
            </Text>
          </TouchableOpacity>
        </View>
      </ScrollView>
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
  card: {
    padding: 20,
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
    margin: 20,
    gap: 20,
    width: '90%',
  },
  buttonContainer: {
    flexDirection: 'row',
    justifyContent: 'center',
    gap: 5,
    flexWrap: 'wrap',
  },
  textContainer: {
    gap: 20,
  },
  bullet: {
    fontSize: 20,
    marginRight: 10,
  },
  subtitle: {
    fontSize: 16,
    fontWeight: '600',
  },
});
