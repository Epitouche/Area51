import { NavigationContainer } from '@react-navigation/native';
import AppNavigator from './navigation/AppNavigator';
import AppProvider from './context/AppContext';

export default function App() {
  return (
    <NavigationContainer>
      <AppProvider>
        <AppNavigator />
      </AppProvider>
    </NavigationContainer>
  );
}
