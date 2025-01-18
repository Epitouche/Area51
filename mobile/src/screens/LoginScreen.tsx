import React, { useContext, useState } from 'react';
import {
  View,
  TextInput,
  TouchableOpacity,
  Text,
  StyleSheet,
  Image,
} from 'react-native';
import { AppContext } from '../context/AppContext';
import { loginApiCall } from '../service';
import { AboutJson, AboutJsonParse, LoginProps } from '../types';
import { OauthLoginButton, IpInput } from '../components';
import { globalStyles } from '../styles/global_style';

interface LoginFunctionProps {
  serverIp: string;
  setIsConnected: (isConnected: boolean) => void;
  isBlackTheme: boolean;
  aboutJson?: AboutJson;
}

interface NoIpProps {
  isBlackTheme?: boolean;
  setAboutJson: (aboutJson: AboutJson) => void;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
  aboutJson: AboutJson | undefined;
  setServerIp: (serverIp: string) => void;
  serverIp: string;
}

function Login({
  isBlackTheme,
  serverIp,
  setIsConnected,
  aboutJson,
}: LoginFunctionProps) {
  const [forms, setForms] = useState<LoginProps>({
    username: '',
    password: '',
  });
  const [message, setMessage] = useState('');

  const handleLogin = async () => {
    const loginSuccessful = await loginApiCall({
      apiEndpoint: serverIp,
      formsLogin: forms,
      setMessage,
    });
    if (loginSuccessful) {
      setIsConnected(true);
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
          accessibilityLabel="Login">
          LOG IN
        </Text>
        <View style={styles.inputBox}>
          <TextInput
            style={[
              isBlackTheme ? globalStyles.inputBlack : globalStyles.input,
              { width: '90%' },
            ]}
            autoCapitalize="none"
            placeholder="Username"
            value={forms.username}
            onChangeText={username => setForms({ ...forms, username })}
            accessibilityLabel="Username"
          />
          <TextInput
            style={[
              isBlackTheme ? globalStyles.inputBlack : globalStyles.input,
              { width: '90%' },
            ]}
            secureTextEntry
            value={forms.password}
            placeholder="Password"
            onChangeText={password => setForms({ ...forms, password })}
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
        <View style={{ width: '90%' }}>
          <TouchableOpacity
            style={[globalStyles.buttonFormat, globalStyles.terciaryLight]}
            onPress={handleLogin}>
            <Text
              style={[globalStyles.textColorBlack, globalStyles.textFormat]}
              accessibilityLabel="Login Button">
              Login
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

export default function LoginScreen() {
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
    <Login
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
    gap: 40,
  },
  header: {
    fontSize: 32,
    color: '#fff',
    fontWeight: 'bold',
    marginTop: '20%',
  },
  inputBox: {
    width: '100%',
    alignItems: 'center',
    gap: 30,
    marginTop: '10%',
  },
  input: {
    width: '90%',
    borderBottomWidth: 1,
    borderColor: '#F7FAFB',
    padding: 5,
    marginVertical: 10,
    fontSize: 16,
    color: 'white',
  },
  loginButton: {
    width: '35%',
    backgroundColor: '#F7FAFB',
    justifyContent: 'center',
    alignItems: 'center',
  },
  passwordText: { color: 'white' },
  forgotPassword: {
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
});
