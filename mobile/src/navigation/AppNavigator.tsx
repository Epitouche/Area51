import React, { useContext } from 'react';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { AppContext } from '../context/AppContext';
import { AuthParamList, AppStackList} from '../types';

// Screens
import LoginScreen from '../screens/LoginScreen';
import RegisterScreen from '../screens/RegisterScreen';
import WorkflowScreen from '../screens/WorkflowScreen';
import HomeScreen from '../screens/HomeScreen';
import ServiceScreen from '../screens/ServiceScreen';
import ActionOrReactionScreen from '../screens/ActionOrReactionScreen';

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
      <AppStack.Screen name="ActionOrReaction" component={ActionOrReactionScreen} />
    </AppStack.Navigator>
  );
}

function AuthStackScreen() {
  return (
    <AuthStack.Navigator screenOptions={{ headerShown: false }}>
      <AuthStack.Screen name="Home" component={HomeScreen} />
      <AuthStack.Screen name="Login" component={LoginScreen} />
      <AuthStack.Screen name="Register" component={RegisterScreen} />
    </AuthStack.Navigator>
  );
}

function OptionsStackScreen() {
  return (
    <Tab.Navigator screenOptions={{ headerShown: false }}>
      <Tab.Screen name="Home" component={HomeScreen} />
      <Tab.Screen name="Workflows" component={WorkflowScreen} />
      <Tab.Screen name="Service" component={ServiceScreen} />
    </Tab.Navigator>
  );
}

// Main App
export default function App() {
  const { isConnected } = useContext(AppContext);

  return isConnected ? <AppStackScreen /> : <AuthStackScreen />;
}
