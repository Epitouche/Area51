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
import LinearGradient from 'react-native-linear-gradient';
import { RegisterProps } from '../types';
import { registerApiCall } from '../service/auth';
import { AppContext } from '../context/AppContext';
import { RootStackParamList } from '../navigation/AppNavigator';
import { StackNavigationProp } from '@react-navigation/stack';


type RegisterScreenNavigationProp = StackNavigationProp<
  RootStackParamList,
  'Register'
>;

type Props = {
  navigation: RegisterScreenNavigationProp;
};

export default function RegisterScreen({ navigation }: Props) {
  const [message, setMessage] = useState('');
  const [forms, setForms] = useState<RegisterProps>({
    email: '',
    password: '',
    username: '',
  });
  const { serverIp } = useContext(AppContext);

  const handleLogin = async () => {
    setMessage('');
    if (await registerApiCall({
      apiEndpoint: serverIp,
      formsRegister: forms,
      setMessage,
    }))
      navigation.navigate('Dashboard');
  };

  return (
    <LinearGradient colors={['#7874FD', '#B225EE']} style={styles.container}>
      <Text style={styles.header}>REGISTER IN</Text>
      <View style={styles.inputBox}>
        <TextInput
          style={styles.input}
          placeholder="Username"
          keyboardType="default"
          autoCapitalize="none"
          value={forms.username}
          onChangeText={text => setForms({ ...forms, username: text })}
        />
        <TextInput
          style={styles.input}
          placeholder="Email"
          keyboardType="email-address"
          autoCapitalize="none"
          value={forms.email}
          onChangeText={text => setForms({ ...forms, email: text })}
        />
        <TextInput
          style={styles.input}
          secureTextEntry
          placeholder="Password"
          value={forms.password}
          onChangeText={text => setForms({ ...forms, password: text })}
          autoCapitalize="none"
        />
      </View>
      {message !== '' && <Text style={styles.errorMessage}>{message}</Text>}
      {/* <View style={styles.checkboxContainer}>
        <TouchableOpacity
          onPress={() => setRememberMe(!rememberMe)}
          style={styles.checkbox}>
          {rememberMe ? (
            <Text style={styles.checkboxText}>✔</Text>
          ) : (
            <Text style={styles.checkboxText}>☐</Text>
          )}
        </TouchableOpacity>
        <Text style={styles.rememberMeText}>Remember me</Text>
      </View> */}
      <Button mode="contained" style={styles.loginButton} onPress={handleLogin}>
        <Text style={styles.text}>Login</Text>
      </Button>
      <View style={styles.line} />
      <View style={styles.socialButtonBox}>
        <Button style={styles.button}>
          <View style={styles.buttonContent}>
            <Image
              source={{
                uri: 'https://img.icons8.com/?size=100&id=12599&format=png',
              }}
              style={styles.icon}
            />
            <Text style={styles.text}>Github</Text>
          </View>
        </Button>
        <Button style={styles.button}>
          <View style={styles.buttonContent}>
            <Image
              source={{
                uri: 'https://img.icons8.com/color/48/google-logo.png',
              }}
              style={styles.icon}
            />
            <Text style={styles.text}>Google</Text>
          </View>
        </Button>
      </View>
      <View style={styles.forgotPasswordBox}>
        <TouchableOpacity>
          <Text style={styles.forgotPassword}>Forgot Password?</Text>
        </TouchableOpacity>
      </View>
    </LinearGradient>
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
  // checkboxContainer: {
  //   flexDirection: 'row',
  //   alignItems: 'center',
  //   marginBottom: 20,
  // },
  // checkbox: {
  //   width: 24,
  //   height: 24,
  //   borderWidth: 2,
  //   borderColor: '#fff',
  //   borderRadius: 20,
  //   justifyContent: 'center',
  //   alignItems: 'center',
  //   marginRight: 10,
  // },
  // checkboxText: {
  //   fontSize: 18,
  //   color: '#fff',
  // },
  // rememberMeText: {
  //   color: '#fff',
  //   fontSize: 16,
  // },

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
  loginButton: {
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
  errorMessage: {
    color: 'red',
    fontSize: 16,
    marginTop: 10,
  },
});
