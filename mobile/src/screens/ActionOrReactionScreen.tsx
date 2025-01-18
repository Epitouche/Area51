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
import { AppStackList, ServicesParse, OptionValues } from '../types';
import { RouteProp, useNavigation, useRoute } from '@react-navigation/native';

interface Values {
  id: number;
  name: string;
  description: string;
  options: OptionValues[] | null;
}

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
  const { isAction, setValues } = route.params;
  const { servicesConnected, isBlackTheme } = useContext(AppContext);

  useEffect(() => {
    setSelectedActionOrReactionId(undefined);
  }, [selectedService]);

  const renderOptionFields = (
    options: OptionValues[],
    isBlackTheme: boolean,
  ) => {
    return options.map((option: OptionValues, index: number) => {
      if (
        typeof option.var === 'object' &&
        !Array.isArray(option.var) &&
        option.value !== null
      ) {
        return (
          <View key={index}>
            <Text style={[globalStyles.textColor, globalStyles.textFormat]}>
              {option.name.charAt(0).toUpperCase() + option.name.slice(1)}
            </Text>
            <View style={{ marginLeft: 20 }}>
              {renderOptionFields(
                Object.entries(option.var).map(([name, value]) => ({
                  name,
                  value: '',
                  var: value,
                })),
                isBlackTheme,
              )}
            </View>
          </View>
        );
      }
      return (
        <>
          <Text style={[globalStyles.textColor, globalStyles.textFormat]}>
            {option.name.charAt(0).toUpperCase() + option.name.slice(1)}
          </Text>
          <TextInput
            key={index}
            placeholder={`Ex: ${option.var}`}
            defaultValue={String(option.value)}
            accessibilityLabel={`Enter the Options for ${option.name} de type ${option.var}`}
            style={[
              isBlackTheme ? globalStyles.input : globalStyles.inputBlack,
            ]}
            onChangeText={text => {
              const updatedOptions =
                options.map((opt: OptionValues, idx: number) => {
                  if (idx === index) {
                    return { ...opt, value: text };
                  }
                  return opt;
                }) || [];
              setSelectedActionOrReactionId({
                ...selectedActionOrReactionId,
                options: updatedOptions,
              } as Values);
            }}
          />
        </>
      );
    });
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
                            options:
                              action.options?.map(option => ({
                                ...option,
                                value: '',
                              })) || null,
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
                          options:
                            reaction.options?.map(option => ({
                              ...option,
                              value: '',
                            })) || null,
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
              {selectedActionOrReactionId?.options &&
                selectedActionOrReactionId.options.length > 0 && (
                  <>
                    <View
                      style={{ flexDirection: 'row', alignItems: 'center' }}>
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
                    {renderOptionFields(
                      selectedActionOrReactionId.options,
                      isBlackTheme,
                    )}
                    {/* // return (
                      //   <TextInput
                      //     key={index}
                      //     placeholder={`Enter ${option.name}`}
                      //     defaultValue={String(option.value)}
                      //     accessibilityLabel={
                      //       'Enter the Options for ' +
                      //       option.name +
                      //       ' de type ' +
                      //       option.type
                      //     }
                      //     keyboardType={
                      //       option.type === 'string' ? 'default' : 'numeric'
                      //     }
                      //     style={[
                      //       isBlackTheme
                      //         ? globalStyles.input
                      //         : globalStyles.inputBlack,
                      //     ]}
                      //     onChangeText={text => {
                      //       const updatedOptions =
                      //         selectedActionOrReactionId.options?.map(
                      //           (opt, idx) => {
                      //             if (idx === index) {
                      //               return { ...opt, value: text };
                      //             }
                      //             return opt;
                      //           },
                      //         ) || [];
                      //       setSelectedActionOrReactionId({
                      //         ...selectedActionOrReactionId,
                      //         options: updatedOptions,
                      //       });
                      //     }}
                      //   />
                      // ); */}
                    {/* })} */}
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
                  options: selectedActionOrReactionId.options || [],
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
  saveButton: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 20,
    marginBottom: 40,
    padding: 10,
  },
  subtitle: {
    fontSize: 16,
    fontWeight: '600',
  },
});
