<template>
    <v-simple-table dense>
        <template v-slot:default>
        <thead>
            <tr>
                <th v-for="item in TABLE_HEADERS"  :key="item.text" class="spet-text-left">{{item.text}}</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in SORT_STUDENTS" :key="item.name">
                <td v-for="el in TABLE_HEADERS" :key="el.text">{{ item[el.value] }}</td>
            </tr>
        </tbody>
        </template>
  </v-simple-table>
</template>

<script>
export default {
    name: "Table",
    computed: {
        TABLE_HEADERS() {
            return this.$store.state.students.TABLE_HEADERS
        },
        SORT_STUDENTS() {
            return this.$store.state.students.SORT_STUDENTS
        },
    }
}
</script>

<style>
.spet-text-left {
    text-align: left;
}
</style>