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
import { GithubLoginButton, IpInput } from '../components';
import { globalStyles } from '../styles/global_style';

export default function LoginScreen() {
  const [forms, setForms] = useState<LoginProps>({
    username: '',
    password: '',
  });
  const { serverIp, setIsConnected } = useContext(AppContext);
  const [message, setMessage] = useState('');
  const [token, setToken] = useState('');

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
    const githubLoginSuccessful = await githubLogin(serverIp, setToken);
    if (githubLoginSuccessful) {
      setIsConnected(true);
    }
  };

  return (
    <View style={globalStyles.wallpaper}>
      <View style={globalStyles.container}>
        <Text style={styles.header}>LOG IN</Text>
        {serverIp === '' ? (
          <IpInput />
        ) : (
          <>
            <View style={styles.inputBox}>
              <TextInput
                style={styles.input}
                autoCapitalize="none"
                placeholder="Username"
                value={forms.username}
                onChangeText={username => setForms({ ...forms, username })}
              />
              <TextInput
                style={styles.input}
                secureTextEntry
                value={forms.password}
                placeholder="Password"
                onChangeText={password => setForms({ ...forms, password })}
                autoCapitalize="none"
              />
            </View>
            <Button
              mode="contained"
              style={globalStyles.buttonColor}
              onPress={handleLogin}>
              <Text
                style={[
                  globalStyles.textBlack,
                  { fontSize: 14, fontWeight: 'bold' },
                ]}>
                Login
              </Text>
            </Button>
            <View style={styles.line} />
            <View style={styles.socialButtonBox}>
              <GithubLoginButton handleGithubLogin={handleGithubLogin} />
              <Button style={globalStyles.buttonColor}>
                <View style={styles.buttonContent}>
                  <Image
                    source={{
                      uri: 'https://img.icons8.com/color/48/google-logo.png',
                    }}
                    style={styles.icon}
                  />
                  <Text
                    style={[globalStyles.textBlack, { fontWeight: 'bold' }]}>
                    Google
                  </Text>
                </View>
              </Button>
            </View>
            {message !== '' && (
              <Text style={styles.passwordText}>{message}</Text>
            )}
            <View style={styles.forgotPasswordBox}>
              <TouchableOpacity>
                <Text style={styles.forgotPassword}>Forgot Password?</Text>
              </TouchableOpacity>
            </View>
          </>
        )}
      </View>
    </View>
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
  line: {
    width: '90%',
    height: 2,
    backgroundColor: '#F7FAFB',
    borderRadius: 2,
    marginBottom: 16,
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
