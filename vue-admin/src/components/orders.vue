<template class="orders">
<div>
  <v-dialog v-model="dialog" max-width="500px">
    <v-card>
      <v-card-title>
        <span class="headline">{{ formTitle }}</span>
      </v-card-title>

      <v-card-text>
        <v-container grid-list-md>
          <v-layout wrap>
            <v-flex xs12>
              <v-select
                    v-model="editedItem.status"
                    :items="status"
                    item-text="value"
                    item-value="value"
                    :error-messages="selectErrors"
                    return-object
                    label="Estado"
                    required
                  ></v-select>
            </v-flex>
          </v-layout>
        </v-container>
      </v-card-text>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="error" outline  @click.native="close">Cancelar</v-btn>
        <v-btn color="primary" outline  @click.native="save">Guardar</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
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
        <td >{{ props.item.id }}</td>
        <td >{{ props.item.initial_value }}</td>
        <td >{{ props.item.final_value }}</td>
        <td >{{ props.item.discount }}</td>
        <td >{{ props.item.status }}</td>
        <td >{{ props.item.clients.email }}</td>
        <td >{{ props.item.gateways.name }}</td>
        <td >{{ props.item.countries.name }}</td>
        <td >{{ props.item.payment_id }}</td>
        <td class="justify-center layout px-0">
          <v-icon
            title
            class="mr-2" color="primary"
            @click="editItem(props.item)"
          >
            edit
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
    name: 'orders',
    props: ['search'],
    created () {
      this.$store.dispatch('getAll', {state: this.viewName})
    },
    mounted () {
    },
    data () {
      return {
        status: [
          { value: 'COMPLETED' },
          { value: 'CANCELED' },
          { value: 'WAITING' },
          { value: 'FAILED' }
        ],
        selectErrors: [],
        pagination: {},
        dialog: false,
        editedIndex: -1,
        viewNameESP: 'Ordenes'
      }
    },
    methods: {
      editItem (item) {
        this.editedIndex = this.list.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog = true
      },
      deleteItem (item) {
        item.index = this.list.indexOf(item)
        let params = {
          state: this.viewName,
          item: item
        }

        confirm('Esta seguro que desea eliminar este elemento?') && this.$store.dispatch('deleteOne', params)
      },
      close () {
        this.dialog = false
        setTimeout(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        }, 300)
      },
      save () {
        let params = {
          state: this.viewName,
          item: this.editedItem
        }

        if (this.editedIndex > -1) {
          this.$store.dispatch('updateOne', params)
        } else {
          this.$store.dispatch('create', params)
        }
        this.close()
      }
    },
    watch: {
      dialog (val) {
        val || this.close()
      }
    },
    computed: {
      headers () {
        return this.$store.state[this.viewName].struct
      },
      list () {
        return this.$store.getters.getAll('orders')
      },
      pages () {
        if (this.pagination.rowsPerPage == null ||
          this.pagination.totalItems == null
        ) return 0

        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      },
      viewName () {
        return this.$route.name
      },
      formTitle () {
        var title = this.editedIndex === -1 ? 'Nuevo ' : 'Editar '
        return title + this.viewNameESP
      },
      defaultItem () {
        return this.$store.state[this.viewName].defaultItem
      },
      editedItem: {
        get () {
          return this.$store.state[this.viewName].editedItem
        },
        set (value) {
          this.$store.state[this.viewName].editedItem = value
          return this.$store.state[this.viewName].editedItem
        }
      }
    }
}
</script>