import { StyleSheet } from 'react-native';

export const globalStyles = StyleSheet.create({
  ipBox: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 20,
  },
  container: {
    width: '100%',
    padding: '1%',
    paddingTop: '20%',
    alignItems: 'center',
    gap: 20,
  },
  line: {
    width: '90%',
    height: 2,
    borderRadius: 2,
    marginBottom: 16,
  },
  textFormat: {
    fontSize: 16,
    fontWeight: 'bold',
  },
  buttonFormat: {
    width: 'auto',
    padding: 10,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
  },

  // White Theme
  wallpaper: {
    flex: 1,
    alignItems: 'center',
    backgroundColor: '#F7FAFB',
  },
  textColor: {
    color: '#0a0a0a',
  },
  title: {
    fontSize: 30,
    color: '#0a0a0a',
    fontWeight: 'bold',
  },
  subtitle: {
    fontSize: 20,
    color: '#0a0a0a',
    fontWeight: 'semibold',
  },
  input: {
    borderBottomWidth: 1,
    borderColor: '#1A1A1A',
    padding: 5,
    marginVertical: 10,
    fontSize: 16,
    color: '#0a0a0a',
  },
  lineColor: {
    backgroundColor: '#1A1A1A',
  },

  // Black Theme
  wallpaperBlack: {
    flex: 1,
    alignItems: 'center',
    backgroundColor: '#222831',
  },
  textColorBlack: {
    color: '#f5f5f5',
  },
  titleBlack: {
    fontSize: 30,
    color: '#f5f5f5',
    fontWeight: 'bold',
  },
  subtitleBlack: {
    fontSize: 20,
    color: '#f5f5f5',
    fontWeight: 'semibold',
  },
  inputBlack: {
    borderBottomWidth: 1,
    borderColor: '#F7FAFB',
    padding: 5,
    marginVertical: 10,
    fontSize: 16,
    color: '#f5f5f5',
  },
  lineColorBlack: {
    backgroundColor: '#F7FAFB',
  },

  //color
  primaryDark: {
    backgroundColor: '#1A1A1A',
  },
  secondaryDark: {
    backgroundColor: '#222831',
  },
  secondaryDark400: {
    backgroundColor: '#4e535a',
  },
  terciaryDark: {
    backgroundColor: '#550195',
  },

  primaryLight: {
    backgroundColor: '#E8E9E9',
  },
  secondaryLight: {
    backgroundColor: '#F7FAFB',
  },
  terciaryLight: {
    backgroundColor: '#8d01f9',
  },
});
