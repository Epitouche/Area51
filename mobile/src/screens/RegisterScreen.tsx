import React, { useContext, useState } from 'react';
import {
  View,
  TextInput,
  TouchableOpacity,
  Text,
  StyleSheet,
} from 'react-native';
import { AboutJson, AuthParamList, NoIpProps, RegisterProps } from '../types';
import { registerApiCall } from '../service/auth';
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
            placeholderTextColor={isBlackTheme ? '#0a0a0a' : 'f5f5f5'}
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
            placeholderTextColor={isBlackTheme ? '#0a0a0a' : 'f5f5f5'}
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
            placeholderTextColor={isBlackTheme ? '#0a0a0a' : 'f5f5f5'}
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
  container: {
    flex: 1,
    alignItems: 'center',
    padding: '3%',
    gap: 30,
  },
  header: {
    fontSize: 32,
    color: '#fff',
    fontWeight: 'bold',
    marginTop: '20%',
  },

  // Input Section
  inputBox: {
    width: '100%',
    alignItems: 'center',
    gap: 30,
    marginTop: '10%',
  },
  input: {
    width: '100%',
    borderBottomWidth: 1,
    borderColor: '#F7FAFB',
    padding: 5,
    marginVertical: 10,
    fontSize: 16,
    color: 'white',
  },

  // Button Section
  registerButton: {
    width: '35%',
    backgroundColor: '#F7FAFB',
    justifyContent: 'center',
    alignItems: 'center',
  },
  passwordText: {
    color: 'white',
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
  line: {
    width: '90%',
    height: 2,
    backgroundColor: '#F7FAFB',
    borderRadius: 2,
    marginBottom: 16,
  },

  socialButtonBox: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
  button: {
    width: '50%',
    marginTop: 10,
    marginBottom: 10,
    backgroundColor: '#F7FAFB',
    alignItems: 'center',
    flexDirection: 'row',
    justifyContent: 'center',
  },
  buttonContent: {
    flexDirection: 'row',
    alignItems: 'center',
  },

  text: {
    color: '#5C5C5C',
    fontSize: 16,
    fontWeight: 'bold',
  },
  icon: {
    marginRight: 15,
    width: 25,
    height: 25,
  },
  errorMessage: {
    color: 'red',
    fontSize: 16,
    marginTop: 10,
  },
});
