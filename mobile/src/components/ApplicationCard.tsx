import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { deleteToken, saveToken, deleteUser } from '../service';

interface ApplicationCardProps {
  isBlackTheme?: boolean;
  setIsBlackTheme: (isBlackTheme: boolean) => void;
  setIsConnected: (isConnected: boolean) => void;
  token: string;
  serverIp: string;
}

export function ApplicationCard({
  isBlackTheme,
  setIsBlackTheme,
  setIsConnected,
  token,
  serverIp,
}: ApplicationCardProps) {
  const handleLogout = () => {
    setIsConnected(false);
    deleteToken('token');
  };

  const handleDeleteAccount = () => {
    setIsConnected(false);
    deleteToken('token');
    deleteUser({ apiEndpoint: serverIp, token });
  };

  const handleTheme = async () => {
    setIsBlackTheme(!isBlackTheme);
    await deleteToken('isBlackTheme');
    await saveToken('isBlackTheme', (!isBlackTheme).toString());
  };

  return (
    <View
      style={[
        isBlackTheme ? globalStyles.primaryLight : globalStyles.terciaryLight,
        styles.card,
      ]}>
      <View style={{ width: '100%', alignItems: 'center' }}>
        <View style={{ gap: 10, width: '100%', marginBottom: 20 }}>
          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <Text
              style={[
                styles.bullet,
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
              ]}
              accessibilityLabel="Bullet">
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}
              accessibilityLabel={'Choose your app theme'}>
              Choose your app theme
            </Text>
          </View>
          <View style={{ width: '100%' }}>
            <TouchableOpacity
              style={[
                globalStyles.buttonFormat,
                isBlackTheme
                  ? globalStyles.secondaryDark
                  : globalStyles.primaryLight,
              ]}
              onPress={handleTheme}>
              <Text
                style={[
                  isBlackTheme
                    ? globalStyles.textColorBlack
                    : globalStyles.textColor,
                  globalStyles.textFormat,
                ]}
                accessibilityLabel={
                  isBlackTheme ? 'Theme Black' : 'Theme White'
                }>
                {isBlackTheme ? 'Theme Black' : 'Theme White'}
              </Text>
            </TouchableOpacity>
          </View>
        </View>

        <View style={{ gap: 10, width: '100%' }}>
          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <Text
              style={[
                styles.bullet,
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
              ]}
              accessibilityLabel="Bullet">
              •
            </Text>
            <Text
              style={[
                isBlackTheme
                  ? globalStyles.textColor
                  : globalStyles.textColorBlack,
                styles.subtitle,
              ]}
              accessibilityLabel="Logout ?">
              Logout ?
            </Text>
          </View>
          <View style={{ width: '100%', gap: 20 }}>
            <TouchableOpacity
              style={[
                globalStyles.buttonFormat,
                { backgroundColor: '#f44336' },
              ]}
              onPress={handleLogout}>
              <Text
                style={[globalStyles.textColorBlack, globalStyles.textFormat]}
                accessibilityLabel="Logout">
                Logout
              </Text>
            </TouchableOpacity>
            {token &&
              (token === 'Error: token not found' || token !=='') && (
                  <TouchableOpacity
                    style={[
                      globalStyles.buttonFormat,
                      { backgroundColor: '#f44336' },
                    ]}
                    onPress={handleDeleteAccount}>
                    <Text
                      style={[
                        globalStyles.textColorBlack,
                        globalStyles.textFormat,
                      ]}
                      accessibilityLabel="Delete Account">
                      Delete Account
                    </Text>
                  </TouchableOpacity>
                )}
          </View>
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    padding: 20,
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
    margin: 20,
  },
  button: {
    width: '100%',
  },
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
  bullet: {
    fontSize: 20,
    marginRight: 10,
  },
  subtitle: {
    fontSize: 16,
    fontWeight: '600',
  },
});
