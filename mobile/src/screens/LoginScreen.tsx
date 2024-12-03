import React, {useState} from 'react';
import {
  View,
  TextInput,
  TouchableOpacity,
  Text,
  StyleSheet,
} from 'react-native';
import {Button} from 'react-native-paper';
import LinearGradient from 'react-native-linear-gradient';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [rememberMe, setRememberMe] = useState(false);

  const handleLogin = () => {
    console.log('Logging in with', email, password);
  };

  return (
    <LinearGradient colors={['#7874FD', '#B225EE']} style={styles.container}>
      <Text style={styles.header}>LOG IN</Text>
      <TextInput
        style={styles.input}
        placeholder="Email"
        keyboardType="email-address"
        value={email}
        onChangeText={setEmail}
      />
      <TextInput
        style={styles.input}
        placeholder="Password"
        secureTextEntry
        value={password}
        onChangeText={setPassword}
      />
      <View style={styles.checkboxContainer}>
        <TouchableOpacity
          onPress={() => setRememberMe(!rememberMe)}
          style={styles.checkbox}>
          {rememberMe ? (
            <Text style={styles.checkboxText}>✔</Text>
          ) : (
            <Text style={styles.checkboxText}>☐</Text>
          )}
        </TouchableOpacity>
        <Text style={styles.rememberMeText}>Remember me</Text>
      </View>
      <Button mode="contained" style={styles.loginButton} onPress={handleLogin}>
        <Text style={styles.socialButtonText}>Login</Text>
      </Button>
      <View style={styles.line} />
      <View style={styles.socialButtons}>
        <Button style={styles.socialButton}>
          <Text style={styles.socialButtonText}>Github</Text>
        </Button>
        <Button style={styles.socialButton}>
          <Text style={styles.socialButtonText}>Google</Text>
        </Button>
      </View>
      <View style={styles.forgotPasswordBox}>
        <TouchableOpacity>
          <Text style={styles.forgotPassword}>Forgot Password?</Text>
        </TouchableOpacity>
      </View>
    </LinearGradient>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  header: {
    fontSize: 32,
    color: '#fff',
    fontWeight: 'bold',
    marginBottom: 30,
  },
  input: {
    width: '100%',
    height: 50,
    backgroundColor: '#F7FAFB',
    marginBottom: 15,
    paddingLeft: 10,
    borderRadius: 20,
    fontSize: 16,
    color: 'black',
  },
  checkboxContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 20,
  },
  checkbox: {
    width: 24,
    height: 24,
    borderWidth: 2,
    borderColor: '#fff',
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 10,
  },
  checkboxText: {
    fontSize: 18,
    color: '#fff',
  },
  rememberMeText: {
    color: '#fff',
    fontSize: 16,
  },
  loginButton: {
    width: '35%',
    borderRadius: 20,
    marginBottom: 20,
    backgroundColor: '#F7FAFB',
    justifyContent: 'center',
    alignItems: 'center',
  },
  socialButtons: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%',
    marginBottom: 20,
  },
  socialButton: {
    width: '45%',
    marginBottom: 10,
    borderRadius: 20,
    backgroundColor: '#F7FAFB',
  },
  forgotPassword: {
    color: '#fff',
    textDecorationLine: 'underline',
    fontSize: 16,
  },
  forgotPasswordBox: {
    width: '100%',
    justifyContent: 'center',
    alignItems: 'center',
  },
  socialButtonText: {
    color: '#5C5C5C',
    fontSize: 16,
    fontWeight: 'bold',
  },
  line: {
    width: '80%',
    height: 2,
    backgroundColor: '#F7FAFB',
    borderRadius: 2,
    marginBottom: 16,
  },
});
