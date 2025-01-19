import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppStackList, Workflow } from '../types';
import { NavigationProp } from '@react-navigation/native';

type WorkflowTabProps = {
  workflows?: Workflow[];
  isBlackTheme?: boolean;
  navigation: NavigationProp<AppStackList>;
};

export function WorkflowTab({
  workflows,
  isBlackTheme,
  navigation,
}: WorkflowTabProps) {
  return (
    <View style={styles.container}>
      <View
        style={[
          styles.headerRow,
          isBlackTheme
            ? globalStyles.secondaryLight
            : globalStyles.secondaryDark,
        ]}>
        <Text
          style={[
            styles.headerCell,
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
            globalStyles.textFormat,
          ]}>
          Name
        </Text>
        <Text
          style={[
            styles.headerCell,
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
            globalStyles.textFormat,
          ]}>
          Is Active
        </Text>
        <Text
          style={[
            styles.headerCell,
            isBlackTheme ? globalStyles.textColor : globalStyles.textColorBlack,
            globalStyles.textFormat,
          ]}>
          Details
        </Text>
      </View>
      {workflows &&
        workflows.map((workflow, index) => (
          <View
            key={index}
            style={[
              styles.row,
              isBlackTheme
                ? globalStyles.secondaryLight
                : globalStyles.secondaryDark,
            ]}>
            <View style={styles.cell}>
              <Text
                style={
                  isBlackTheme
                    ? globalStyles.textColor
                    : globalStyles.textColorBlack
                }>
                {workflow.name}
              </Text>
            </View>
            <View style={styles.cell}>
              <Text
                style={
                  isBlackTheme
                    ? globalStyles.textColor
                    : globalStyles.textColorBlack
                }>
                {workflow.is_active ? 'Yes' : 'No'}
              </Text>
            </View>
            <View style={styles.cell}>
              <TouchableOpacity
                onPress={() =>
                  navigation.navigate('Workflow Details', { workflow })
                }>
                <Text
                  style={
                    isBlackTheme
                      ? globalStyles.textColor
                      : globalStyles.textColorBlack
                  }>
                  ...
                </Text>
              </TouchableOpacity>
            </View>
          </View>
        ))}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    alignContent: 'center',
    justifyContent: 'center',
    width: '90%',
  },
  headerRow: {
    alignContent: 'center',
    flexDirection: 'row',
    borderTopLeftRadius: 5,
    borderTopRightRadius: 5,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  headerCell: {
    width: '33%',
    flex: 1,
    padding: 10,
    fontWeight: 'bold',
    textAlign: 'center',
    color: '#333',
  },
  row: {
    height: 40,
    alignContent: 'center',
    flexDirection: 'row',
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
    justifyContent: 'center',
  },
  cell: {
    justifyContent: 'center',
    alignItems: 'center',
    width: '33%',
  },
  buttonText: {
    color: '#fff',
  },
});
