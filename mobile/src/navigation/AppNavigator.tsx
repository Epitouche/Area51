import React, { useContext } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { AppContext } from '../context/AppContext';
import { RootStackParamList, TabParamList } from '../types';

// Screens
import LoginScreen from '../screens/LoginScreen';
import RegisterScreen from '../screens/RegisterScreen';
import DashboardScreen from '../screens/DashboardScreen';
import HomeScreen from '../screens/HomeScreen';

const Tab = createBottomTabNavigator<TabParamList>();
const AuthStack = createNativeStackNavigator<RootStackParamList>();

function AuthStackScreen() {
  return (
    <AuthStack.Navigator>
      <AuthStack.Screen name="Home" component={HomeScreen} />
      <AuthStack.Screen name="Login" component={LoginScreen} />
      <AuthStack.Screen name="Register" component={RegisterScreen} />
    </AuthStack.Navigator>
  );
}

function AppStackScreen() {
  return (
    <Tab.Navigator>
      <Tab.Screen name="Home" component={HomeScreen} />
      <Tab.Screen name="Dashboard" component={DashboardScreen} />
    </Tab.Navigator>
  );
}

// Main App
export default function App() {
  const { isConnected } = useContext(AppContext);

  return isConnected ? <AppStackScreen /> : <AuthStackScreen />;
}
