import React, { useEffect, useState } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { Button } from 'react-native-paper';
import LinearGradient from 'react-native-linear-gradient';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigation/AppNavigator';
import { deleteToken, checkToken } from '../service/token';

type DashboardNavigationProp = StackNavigationProp<
  RootStackParamList,
  'Dashboard'
>;

type Props = {
  navigation: DashboardNavigationProp;
};
export default function DashboardScreen({ navigation }: Props) {
  const [token, setToken] = useState('');
  const [github, setGithub] = useState('');

  const handleLogout = () => {
    deleteToken('token');
    deleteToken('github');
    navigation.navigate('Home');
  };

  useEffect(() => {
    checkToken('token', setToken);
    checkToken('github', setGithub);
    if (token === 'Error: token not found' && github === 'Error: token not found') {
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
  loginButton: {
    width: '35%',
    backgroundColor: '#F7FAFB',
    justifyContent: 'center',
    alignItems: 'center',
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
  text: {
    color: '#5C5C5C',
    fontSize: 16,
    fontWeight: 'bold',
  },
});
