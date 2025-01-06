import React, { useContext, useState } from 'react';
import {
  View,
  TextInput,
  TouchableOpacity,
  Text,
  StyleSheet,
  Image,
} from 'react-native';
import { Button } from 'react-native-paper';
import { AppContext } from '../context/AppContext';
import { loginApiCall, githubLogin } from '../service';
import { LoginProps } from '../types';
import { OauthLoginButton, IpInput } from '../components';
import { globalStyles } from '../styles/global_style';

interface LoginFunctionProps {
  serverIp: string;
  setIsConnected: Function;
  isBlackTheme: boolean;
}

interface NoIpProps {
  isBlackTheme: boolean;
}

function Login({ isBlackTheme, serverIp, setIsConnected }: LoginFunctionProps) {
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

  const handleGithubLogin = async () => {
    const githubLoginSuccessful = await githubLogin(serverIp);
    if (githubLoginSuccessful) {
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
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
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
          />
        </View>
        <View>
          {message != '' && <Text style={{ color: 'red' }}>{message}</Text>}
        </View>
        <Button
          mode="contained"
          style={styles.loginButton}
          onPress={handleLogin}>
          <Text
            style={[
              isBlackTheme ? globalStyles.textBlack : globalStyles.text,
              { fontSize: 14, fontWeight: 'bold' },
            ]}>
            Login
          </Text>
        </Button>
        <View
          style={[
            globalStyles.line,
            isBlackTheme ? globalStyles.lineColorBlack : globalStyles.lineColor,
          ]}
        />
        <View style={styles.socialButtonBox}>
          <OauthLoginButton
            handleOauthLogin={handleGithubLogin}
            name="Github"
            img="https://img.icons8.com/?size=100&id=12599&format=png"
          />
          <OauthLoginButton
            handleOauthLogin={handleGithubLogin}
            name="Google"
            img="https://img.icons8.com/?size=100&id=12599&format=png"
          />
        </View>
        {message !== '' && <Text style={styles.passwordText}>{message}</Text>}
        <View style={styles.forgotPasswordBox}>
          <TouchableOpacity>
            <Text style={styles.forgotPassword}>Forgot Password?</Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
}

function NoIp({ isBlackTheme }: NoIpProps) {
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
        <IpInput />
      </View>
    </View>
  );
}

export default function LoginScreen() {
  const { serverIp, setIsConnected, isBlackTheme } = useContext(AppContext);

  return serverIp ? (
    <Login
      serverIp={serverIp}
      setIsConnected={setIsConnected}
      isBlackTheme={isBlackTheme}
    />
  ) : (
    <NoIp isBlackTheme={isBlackTheme} />
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
    width: '80%',
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
