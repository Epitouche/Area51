<script setup lang="ts">

const props = defineProps<{
    bgColor?: string;
    textColor?: string;
    icon?: string;
    iconHeight?: number;
    iconWidth?: number;
    ariaLabel?: string;
}>();

const attrs = useAttrs();

const defaultBgColor = "bg-primaryWhite-500 dark:bg-primaryDark-500";
const defaultTextColor = "text-fontBlack dark:text-fontWhite";
const defaultIcon = 'jam:thunder'
const defaulticonHeight = 6
const defaulticonWidth = 6

const computedBgColor = computed(() => props.bgColor ?? defaultBgColor);
const computedTextColor = computed(() => props.textColor ?? defaultTextColor);
const computedIcon = computed(() => props.icon ?? defaultIcon);
const computedIconHeight = computed(() => props.iconHeight ?? defaulticonHeight);
const computedIconWidth = computed(() => props.iconWidth ?? defaulticonWidth);
</script>

<template>
    <div
    :class="[
        'p-3 rounded-xl shadow-lg transform hover:scale-110 transition-transform duration-200',
        computedBgColor,
        computedTextColor,
    ]" v-bind="attrs">
        <Icon
        :name="computedIcon"
        :aria-label="props.ariaLabel"
        :class="[
            'mx-1 mt-1',
            `h-${computedIconHeight}`,
            `w-${computedIconWidth}`
        ]" />
    </div>
</template>
