import React, { useContext, useState } from 'react';
import {
  View,
  TextInput,
  TouchableOpacity,
  Text,
  StyleSheet,
} from 'react-native';
import { AuthParamList, RegisterProps } from '../types';
import { registerApiCall } from '../service/auth';
import { AppContext } from '../context/AppContext';
import { OauthLoginButton, IpInput } from '../components';
import { globalStyles } from '../styles/global_style';
import { NavigationProp, useNavigation } from '@react-navigation/native';

export default function RegisterScreen() {
  const nav = useNavigation<NavigationProp<AuthParamList>>();
  const [message, setMessage] = useState('');
  const [forms, setForms] = useState<RegisterProps>({
    email: '',
    password: '',
    username: '',
  });
  const { serverIp, setIsConnected, isBlackTheme, aboutJson } =
    useContext(AppContext);

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
        {serverIp === '' ? (
          <IpInput />
        ) : (
          <>
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
                onChangeText={text => setForms({ ...forms, password: text })}
                autoCapitalize="none"
                accessibilityLabel="Password"
              />
            </View>
            <View style={{ width: '90%', marginTop: 20 }}>
              <TouchableOpacity
                style={[
                  globalStyles.buttonFormat,
                  isBlackTheme
                    ? globalStyles.terciaryDark
                    : globalStyles.terciaryLight,
                ]}
                onPress={handleRegister}>
                <Text
                  style={[
                    isBlackTheme
                      ? globalStyles.textColorBlack
                      : globalStyles.textColor,
                    globalStyles.textFormat,
                  ]}
                  accessibilityLabel="Register Button">
                  Register
                </Text>
              </TouchableOpacity>
            </View>
            <View style={styles.line} />
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
                  style={styles.forgotPassword}
                  accessibilityLabel="Forgot Password">
                  Forgot Password?
                </Text>
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
