import React, { useContext, useState } from 'react';
import {
  View,
  TextInput,
  TouchableOpacity,
  Text,
  StyleSheet,
} from 'react-native';
import { AboutJson, AuthParamList, NoIpProps, RegisterProps } from '../types';
import { registerApiCall } from '../service';
import { AppContext } from '../context/AppContext';
import { OauthLoginButton, IpInput } from '../components';
import { globalStyles } from '../styles/global_style';
import { NavigationProp, useNavigation } from '@react-navigation/native';

interface RegisterFunctionProps {
  serverIp: string;
  setIsConnected: (isConnected: boolean) => void;
  isBlackTheme: boolean;
  aboutJson?: AboutJson;
}

export function Register({
  isBlackTheme,
  serverIp,
  setIsConnected,
  aboutJson,
}: RegisterFunctionProps) {
  const nav = useNavigation<NavigationProp<AuthParamList>>();
  const [message, setMessage] = useState('');
  const [forms, setForms] = useState<RegisterProps>({
    email: '',
    password: '',
    username: '',
  });

  const handleRegister = async () => {
    setMessage('');
    if (
      await registerApiCall({
        apiEndpoint: serverIp,
        formsRegister: forms,
        setMessage,
      })
    ) {
      setForms({ email: '', password: '', username: '' });
      nav.navigate('Login');
    }
  };

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}
          accessibilityLabel="Register">
          REGISTER IN
        </Text>
        <View style={styles.inputBox}>
          <TextInput
            style={[
              isBlackTheme ? globalStyles.inputBlack : globalStyles.input,
              { width: '90%' },
            ]}
            placeholder="Username"
            keyboardType="default"
            autoCapitalize="none"
            value={forms.username}
            placeholderTextColor={isBlackTheme ? '#f5f5f5': '#0a0a0a'}
            onChangeText={text => setForms({ ...forms, username: text })}
            accessibilityLabel="Username"
          />
          <TextInput
            style={[
              isBlackTheme ? globalStyles.inputBlack : globalStyles.input,
              { width: '90%' },
            ]}
            placeholder="Email"
            keyboardType="email-address"
            autoCapitalize="none"
            value={forms.email}
            placeholderTextColor={isBlackTheme ? '#f5f5f5': '#0a0a0a'}
            onChangeText={text => setForms({ ...forms, email: text })}
            accessibilityLabel="Email"
          />
          <TextInput
            style={[
              isBlackTheme ? globalStyles.inputBlack : globalStyles.input,
              { width: '90%' },
            ]}
            secureTextEntry
            placeholder="Password"
            value={forms.password}
            placeholderTextColor={isBlackTheme ? '#f5f5f5': '#0a0a0a'}
            onChangeText={text => setForms({ ...forms, password: text })}
            autoCapitalize="none"
            accessibilityLabel="Password"
          />
        </View>
        <View>
          {message != '' && (
            <Text style={{ color: 'red' }} accessibilityLabel="Error Message">
              {message}
            </Text>
          )}
        </View>
        <View style={{ width: '90%', marginTop: 20 }}>
          <TouchableOpacity
            style={[globalStyles.buttonFormat, globalStyles.terciaryLight]}
            onPress={handleRegister}>
            <Text
              style={[globalStyles.textColorBlack, globalStyles.textFormat]}
              accessibilityLabel="Register Button">
              Register
            </Text>
          </TouchableOpacity>
        </View>
        <View
          style={[
            globalStyles.line,
            isBlackTheme ? globalStyles.lineColorBlack : globalStyles.lineColor,
          ]}
        />
        <View style={styles.socialButtonBox}>
          {aboutJson &&
            aboutJson.server.services.map((service, index) => {
              if (!service.is_oauth) return null;
              return (
                <OauthLoginButton
                  key={index}
                  serverIp={serverIp}
                  setIsConnected={setIsConnected}
                  name={service.name}
                  img={service.image}
                  isBlackTheme={isBlackTheme}
                />
              );
            })}
        </View>
        <View style={styles.forgotPasswordBox}>
          <TouchableOpacity>
            <Text
              style={[
                styles.forgotPassword,
                isBlackTheme
                  ? globalStyles.textColorBlack
                  : globalStyles.textColor,
              ]}
              accessibilityLabel="Forgot Password">
              Forgot Password?
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
}

function NoIp({
  isBlackTheme,
  serverIp,
  setServerIp,
  aboutJson,
  setAboutJson,
  setServicesConnected,
}: NoIpProps) {
  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
          LOG IN
        </Text>
        <IpInput
          aboutJson={aboutJson}
          serverIp={serverIp}
          setServerIp={setServerIp}
          setAboutJson={setAboutJson}
          setServicesConnected={setServicesConnected}
          isBlackTheme={isBlackTheme}
        />
      </View>
    </View>
  );
}

export default function RegisterScreen() {
  const {
    serverIp,
    setIsConnected,
    isBlackTheme,
    aboutJson,
    setAboutJson,
    setServerIp,
    setServicesConnected,
  } = useContext(AppContext);

  return serverIp ? (
    <Register
      serverIp={serverIp}
      setIsConnected={setIsConnected}
      isBlackTheme={isBlackTheme}
      aboutJson={aboutJson}
    />
  ) : (
    <NoIp
      isBlackTheme={isBlackTheme}
      aboutJson={aboutJson}
      serverIp={serverIp}
      setAboutJson={setAboutJson}
      setServerIp={setServerIp}
      setServicesConnected={setServicesConnected}
    />
  );
}

const styles = StyleSheet.create({
  inputBox: {
    width: '100%',
    alignItems: 'center',
    gap: 30,
    marginTop: '10%',
  },
  forgotPassword: {
    color: '#fff',
    textDecorationLine: 'underline',
    fontSize: 16,
  },
  forgotPasswordBox: {
    width: '100%',
    margin: 10,
    justifyContent: 'center',
    alignItems: 'center',
  },

  socialButtonBox: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
});
