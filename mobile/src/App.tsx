import { NavigationContainer } from '@react-navigation/native';
import AppNavigator from './navigation/AppNavigator';
import AppProvider from './context/AppContext';
// import { GestureHandlerRootView } from 'react-native-gesture-handler';
// import { SafeAreaProvider } from 'react-native-safe-area-context';

export default function App() {
  return (
    <NavigationContainer>
      <AppProvider>
        <AppNavigator />
      </AppProvider>
    </NavigationContainer>
  );
}

//  <SafeAreaProvider>
//    <GestureHandlerRootView style={{ flex: 1 }}>
//      <NavigationContainer>
//        <AppProvider>
//          <AppNavigator />
//        </AppProvider>
//      </NavigationContainer>
//    </GestureHandlerRootView>
//  </SafeAreaProvider>;
