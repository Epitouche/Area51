import React, { useContext } from 'react';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { AppContext } from '../context/AppContext';
import { AuthParamList, AppStackList } from '../types';

import {
  HomeIcon,
  SettingIcon,
  ServiceIcon,
  WorkflowIcon,
  LoginIcon,
  RegisterIcon,
} from '../icons/Icons';

import LoginScreen from '../screens/LoginScreen';
import RegisterScreen from '../screens/RegisterScreen';
import WorkflowScreen from '../screens/WorkflowScreen';
import HomeScreen from '../screens/HomeScreen';
import ServiceScreen from '../screens/ServiceScreen';
import ActionOrReactionScreen from '../screens/ActionOrReactionScreen';
import WorkflowDetailsScreen from '../screens/WorkflowDetailsScreen';
import SettingScreen from '../screens/SettingScreen';

const Tab = createBottomTabNavigator<AppStackList>();
const AppStack = createNativeStackNavigator<AppStackList>();
const AuthStack = createBottomTabNavigator<AuthParamList>();

function AppStackScreen() {
  return (
    <AppStack.Navigator>
      <AppStack.Screen
        options={{ headerShown: false }}
        name="App"
        component={OptionsStackScreen}
      />
      <AppStack.Screen name="Options" component={ActionOrReactionScreen} />
      <AppStack.Screen
        name="Workflow Details"
        component={WorkflowDetailsScreen}
      />
    </AppStack.Navigator>
  );
}

function AuthStackScreen() {
  return (
    <AuthStack.Navigator screenOptions={{ headerShown: false }}>
      <AuthStack.Screen
        name="Home"
        component={HomeScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <HomeIcon width={24} height={24} fill={color} />
          ),
        }}
      />
      <AuthStack.Screen
        name="Login"
        component={LoginScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <LoginIcon width={24} height={24} fill={color} />
          ),
        }}
      />
      <AuthStack.Screen
        name="Register"
        component={RegisterScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <RegisterIcon width={24} height={24} fill={color} />
          ),
        }}
      />
    </AuthStack.Navigator>
  );
}

function OptionsStackScreen() {
  return (
    <Tab.Navigator screenOptions={{ headerShown: false }}>
      <Tab.Screen
        name="Home"
        component={HomeScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <HomeIcon width={24} height={24} fill={color} />
          ),
        }}
      />
      <Tab.Screen
        name="Workflows"
        component={WorkflowScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <WorkflowIcon width={24} height={24} fill={color} />
          ),
        }}
      />
      <Tab.Screen
        name="Service"
        component={ServiceScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <ServiceIcon width={24} height={24} fill={color} />
          ),
        }}
      />
      <Tab.Screen
        name="Setting"
        component={SettingScreen}
        options={{
          tabBarIcon: ({ color }) => (
            <SettingIcon width={24} height={24} fill={color} />
          ),
        }}
      />
    </Tab.Navigator>
  );
}

export default function App() {
  const { isConnected } = useContext(AppContext);

  return isConnected ? <AppStackScreen /> : <AuthStackScreen />;
}
