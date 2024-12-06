import React, { useContext, useEffect, useState } from 'react';
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
import { AppContext } from '../context/AppContext';
import { LoginProps } from '../types';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import { deleteToken, checkToken  } from '../service/token';

type DashboardNavigationProp = StackNavigationProp<
  RootStackParamList,
  'Dashboard'
>;

type Props = {
  navigation: DashboardNavigationProp;
};
export default function DashboardScreen({ navigation }: Props) {
  const [forms, setForms] = useState<LoginProps>({
    username: '',
    password: '',
  });
  const { serverIp } = useContext(AppContext);
  const [token, setToken] = useState('');

  const handleLogout = () => {
    deleteToken('token');
    navigation.navigate('Home');
  };

  useEffect(() => {
    checkToken('token', setToken);
    if (token === 'Error: token not found') {
      console.log('Token not found', token);
      navigation.navigate('Home');
    }
  }, [token]);


  return (
    <LinearGradient colors={['#7874FD', '#B225EE']} style={styles.container}>
      <Text style={styles.header}>Dashboard</Text>
      <View style={styles.button}>
        <Button
          mode="contained"
          onPress={handleLogout}
          style={styles.loginButton}>
          <Text style={styles.text}>Logout</Text>
        </Button>
      </View>
    </LinearGradient>
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
});
