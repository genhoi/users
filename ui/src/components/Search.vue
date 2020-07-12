<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-card
                color="primary lighten-2"
                dark
        >
          <v-card-text>
            <v-autocomplete
                    v-model="model"
                    :items="entries"
                    :loading="isLoading"
                    :search-input.sync="search"
                    color="white"
                    hide-no-data
                    hide-selected
                    item-text="full_name"
                    item-value="id"
                    label="Поиск пользователей"
                    placeholder="Начните ввод для поиска"
                    prepend-icon="mdi-database-search"
                    no-filter
                    return-object
            ></v-autocomplete>
          </v-card-text>
        </v-card>


        <v-divider></v-divider>

      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  const axios = require('axios');

  export default {
    name: 'Search',

    data: () => ({
      descriptionLimit: 60,
      entries: [],
      isLoading: false,
      model: null,
      search: null,
    }),
    computed: {
      items () {
        return this.entries
      },
    },
    watch: {
      search (val) {
        // Items have already been requested
        if (this.isLoading) return

        this.isLoading = true

        axios({
          method: 'get',
          url: '/search',
          params: {
            query: val
          }
        }).then(res => {
          this.entries = res.data
        }).finally(() => (this.isLoading = false))
      },
    },
  }
</script>
