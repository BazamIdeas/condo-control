<template class="trashed">
<div>
  <v-toolbar flat color="white">
      <v-toolbar-title class="text-capitalize">{{ viewNameESP }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-title class="text-capitalize">{{ model }}</v-toolbar-title>
  </v-toolbar>
  <v-container fluid>
    <v-card>
      <v-card-title>
        <v-layout row wrap>
          <v-flex offset-xs8 xs4>
              <v-text-field
                v-model="search"
                append-icon="search"
                label="Buscar:"
                single-line
                hide-details
              ></v-text-field>
          </v-flex>
        </v-layout>
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="list"
        class="elevation-1"
        no-data-text="No hubo resultados"
        :search="search"
      >
      <template slot="items" slot-scope="props">
        <td v-for="key in getKeysArray(props.item)" :key="key">
          {{props.item[key]}}
        </td>
        <td class="justify-center layout px-0">
          <v-icon
            title
            class="mr-2" color="primary"
            @click="editItem(props.item)"
          >
            restore
          </v-icon>
          <v-icon
            title
            @click="deleteItem(props.item)" color="error"
          >
            delete
          </v-icon>
        </td>
      </template>
    </v-data-table>
    </v-card>
  </v-container>
</div>
</template>

<script lang="js">
  export default {
    name: 'trashed',
    props: ['model', 'search'],
    created () {
      this.$store.dispatch('getAllTrashed', {state: this.model})
    },
    mounted () {
    },
    data () {
      return {
        pagination: {},
        viewNameESP: 'Papelera'
      }
    },
    methods: {
      deleteItem (item) {
        item.index = this.list.indexOf(item)
        let params = {
          state: this.viewName,
          item: item
        }

        confirm('Esta seguro que desea eliminar este elemento?') && this.$store.dispatch('trash', params)
      },
      restoreItem (item) {
        item.index = this.list.indexOf(item)
        let params = {
          state: this.model,
          item: item
        }

        confirm('Esta seguro que desea restaurar este elemento?') && this.$store.dispatch('restore', params)
      },
      getKeysArray (obj) {
        let keys = []

        for (let i = 0; i < this.headers.length; i++) {
          const element = this.headers[i];
          if(element.text != 'Acciones'){
            keys.push(element.value)
          }
        }

        for (let i = 0; i < keys.length; i++) {
          const key = keys[i]
          if (typeof obj[key] === 'object') {
            if (obj[key].name !== undefined) {
              obj[key] = obj[key].name
            } else if (obj[key].iso !== undefined) {
              obj[key] = obj[key].iso
            } else if (obj[key].id !== undefined) {
              obj[key] = obj[key].id
            } else if (obj[key].title !== undefined) {
              obj[key] = obj[key].title
            }
            continue
          }

          if (obj[key][0] !== undefined) {
            if (obj[key][0].name !== undefined) {
              obj[key][0] = obj[key].name
            } else if (obj[key].iso !== undefined) {
              obj[key][0] = obj[key].iso
            } else if (obj[key].title !== undefined) {
              obj[key][0] = obj[key].title
            }
          }
        }

        console.log(keys)

        return keys
      }
    },
    watch: {
      dialog (val) {
        val || this.close()
      }
    },
    computed: {
      headers () {
        return this.$store.state[this.model].struct
      },
      list () {
        return this.$store.getters.getAllTrashed(this.model)
      },
      pages () {
        if (this.pagination.rowsPerPage == null ||
          this.pagination.totalItems == null
        ) return 0

        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      },
      viewName () {
        return this.$route.name
      }
    }
}
</script>