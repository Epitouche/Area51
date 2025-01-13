<script>
import { onMounted, ref } from 'vue';
import { Chart, registerables } from 'chart.js';

Chart.register(...registerables);

export default {
    setup() {
        const chart = ref(null);
        const timeRanges = ['day', 'week', 'month']
        const selectedRange = ref(timeRanges[0])

        const dayData = {
            labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
            datasets: [{
                label: 'Number of executions',
                data: [12, 19, 3, 5, 2, 3],
                borderWidth: 1
            }]
        };

        const weekData = {
            labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
            datasets: [{
                label: 'Number of executions',
                data: [10, 15, 8, 12, 6, 9, 4],
                borderWidth: 1
            }]
        };

        const monthData = {
            labels: ['Week 1', 'Week 2', 'Week 3', 'Week 4'],
            datasets: [{
                label: 'Number of executions',
                data: [30, 45, 25, 50],
                borderWidth: 1
            }]
        };

        const setData = (period) => {
            selectedRange = period;
            console.log(period);
            switch (period) {
                case 'day':
                    chart.value.data = dayData;
                    break;
                case 'week':
                    chart.value.data = weekData;
                    break;
                case 'month':
                    chart.value.data = monthData;
                    break;
            }
            chart.value.update();
            console.log(period);
        };
        onMounted(() => {
            const ctx = document.getElementById("myChart").getContext('2d');
            chart.value = new Chart(ctx, {
                type: 'bar',
                data: dayData,
                options: {
                    scales: {
                        y: {
                            beginAtZero: true
                        }
                    },
                    responsive: true
                }
            });
        });
        return {
            timeRanges,
            selectedRange,
            setData
        };
    }
};
</script>
<template>
    <div class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6">
        <div class="flex items-center justify-between mb-6">
            <h2 class="text-lg font-semibold text-fontBlack dark:text-fontWhite">Workflow Executions</h2>
            <div class="flex space-x-2">
                <button v-for="time in timeRanges" :key="time" :class="[
                    'px-3 py-1 rounded-lg text-sm font-medium transition-colors',
                    selectedRange === time ? 'bg-tertiary-500 text-white' : 'text-gray-600 dark:text-fontWhite hover:bg-gray-100',
                ]" @click="setData(time)"> {{ time }}
                </button>
            </div>
        </div>
        <div style="position: relative; height: 290px;">
            <canvas id="myChart"></canvas>
        </div>
    </div>
</template>
