import React from 'react';
import AppNavigator from './src/navigation/AppNavigator';
import {AppRegistry} from 'react-native';
import AppProvider from './src/context/AppContext';

export default function App() {
  return (
    <AppProvider>
      <AppNavigator />
    </AppProvider>
  );
}

AppRegistry.registerComponent('Area51', () => App);
