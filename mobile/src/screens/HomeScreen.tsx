// src/screens/HomeScreen.tsx
import React from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigationProp } from '@react-navigation/stack';
import { RootStackParamList } from '../navigation/AppNavigator';

type HomeScreenNavigationProp = StackNavigationProp<RootStackParamList, 'Home'>;

type Props = {
  navigation: HomeScreenNavigationProp;
};

export default function HomeScreen({ navigation }: Props){
  return (
    <View>
      <Text>Home Screen</Text>
      <Button title="Connexion" onPress={() => navigation.navigate('Login')} />
      <Button title="Inscription" onPress={() => navigation.navigate('Register')} />
    </View>
  );
};