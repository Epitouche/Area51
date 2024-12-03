import React from 'react';
import AppNavigator from './src/navigation/AppNavigator';
import { AppRegistry } from 'react-native';

export default function App() {
  return (
      <AppNavigator />
  );
}

AppRegistry.registerComponent('Area51', () => App);