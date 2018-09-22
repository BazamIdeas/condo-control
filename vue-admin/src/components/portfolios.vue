<template class="portfolios">
<div>
  <v-toolbar flat color="white">
      <v-toolbar-title class="text-capitalize">{{ viewNameESP }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-dialog v-model="dialog" max-width="700px">
        <v-btn slot="activator" color="primary" outline class="mb-2">Nuevos {{ viewNameESP }}</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 md6>
                  <v-text-field v-model="editedItem.name" label="Nombre del Portafolio"></v-text-field>
                  <v-text-field v-model="editedItem.client" label="Cliente"></v-text-field>
                  <v-text-field v-model="editedItem.description" label="Descripción"></v-text-field>
                  <v-text-field type="number" v-model="editedItem.priority" label="Prioridad"></v-text-field>
                </v-flex>
                <v-flex xs12 md6>
                  <v-select
                    v-model="editedItem.location"
                    :items="locations"
                    item-text="name"
                    item-value="in"
                    :error-messages="selectErrors"
                    return-object
                    label="Locacion"
                    required
                  ></v-select>
                  <v-select
                    v-model="editedItem.service"
                    :items="services"
                    item-text="name"
                    item-value="in"
                    :error-messages="selectErrors"
                    return-object
                    label="Service"
                    required
                  ></v-select>
                  <v-select
                    v-model="editedItem.activity"
                    :items="activities"
                    item-text="name"
                    item-value="in"
                    :error-messages="selectErrors"
                    return-object
                    label="Actividad"
                    required
                  ></v-select>
                </v-flex>
                <v-flex xs12>
                  <div class="btn btn-primary jbtn-file">Cargar Imagenes: <input
    type="file" v-on:change="fileSelected" multiple></div>
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
        <tr @click="props.expanded = !props.expanded">
          <td>{{ props.item.id }}</td>
          <td >{{ props.item.name }}</td>
          <td >{{ props.item.description }}</td>
          <td >{{ props.item.client }}</td>
          <td >{{ props.item.location.name }}</td>
          <td >{{ props.item.service.name }}</td>
          <td >{{ props.item.activity.name }}</td>
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
        </tr>
      </template>
      <template slot="expand" slot-scope="props">
        <v-container grid-list-md text-xs-center>
          <v-layout row wrap>
            <v-flex
               v-for="img in props.item.images" :key="img.name"
               xs3
            >
              <v-card>
                <v-img
                  :src="urlHosting + img.url"
                  height="200px"
                ></v-img>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn icon @click="imagePriority(img, 1)">
                    <v-icon>favorite</v-icon>
                  </v-btn>
                  <v-btn icon @click="imagePriority(img, 0)">
                    <v-icon>favorite_border</v-icon>
                  </v-btn>
                  <v-btn icon @click="deleteImage(img)">
                    <v-icon>delete</v-icon>
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
      </template>
    </v-data-table>
    </v-card>
  </v-container>
</div>
</template>

<script lang="js">
  export default {
    name: 'portfolios',
    props: ['search'],
    created () {
      this.$store.dispatch('getAll', {state: this.viewName})
      this.$store.dispatch('getAll', {state: 'locations'})
      this.$store.dispatch('getAll', {state: 'services'})
      this.$store.dispatch('getAll', {state: 'activities'})
    },
    mounted () {
    },
    data () {
      return {
        urlHosting: 'http://localhost:9090',
        selectErrors: [],
        pagination: {},
        dialog: false,
        editedIndex: -1,
        viewNameESP: 'Portafolios'
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
          this.$store.dispatch('portfolios/updateOne', params)
        } else {
          this.$store.dispatch('portfolios/create', params)
        }
        this.close()
      },
      imagePriority () {
      },
      deleteImage () {
      },
      fileSelected (e) {
        this.editedItem['files'] = e.target.files
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
        return this.$store.getters.getAll('portfolios')
      },
      locations () {
        return this.$store.getters.getAll('locations')
      },
      services () {
        return this.$store.getters.getAll('services')
      },
      activities () {
        return this.$store.getters.getAll('activities')
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