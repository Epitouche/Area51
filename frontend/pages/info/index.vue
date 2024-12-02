<template>
    <Header/>
    <h1>Info page</h1>
    <article v-for="post in posts" :key="post.id">
        <p>{{ post.title }}</p>
        <p>
            <NuxtLink :to="`/info/${post.id}`">More infos</NuxtLink>
        </p>
    </article>
</template>

<script setup>
const { data: posts } = useFetch('https://jsonplaceholder.typicode.com/posts?_limit=10')
</script>
