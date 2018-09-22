const app = {
  sidemenu:
  [
    {
      icon: 'money',
      title: 'Monedas',
      route: 'currencies'
    },
    {
      icon: 'outlined_flag',
      title: 'Paises',
      route: 'countries'
    },
    {
      icon: 'bookmark_border',
      title: 'Locaciones',
      route: 'locations'
    },
    {
      icon: 'bookmark_border',
      title: 'Sectores',
      route: 'sectors'
    },
    {
      icon: 'local_activity',
      title: 'Actividades',
      route: 'activities'
    },
    {
      icon: 'subtitles',
      title: 'Servicios',
      route: 'services'
    },
    {
      icon: 'description',
      title: 'Ordenes',
      route: 'orders'
    },
    {
      icon: 'payment',
      title: 'Pasarelas',
      route: 'gateways'
    },
    {
      icon: 'card_giftcard',
      title: 'Cupones',
      route: 'coupons'
    },
    /*  {
      icon: 'payment',
      title: 'Metodos de pago',
      route: 'payments-methods'
    }, */
    {
      icon: 'work',
      title: 'Briefs',
      route: 'briefs'
    },
    {
      icon: 'work',
      title: 'Portafolios',
      route: 'portfolios'
    },
    {
      icon: 'face',
      title: 'Clientes',
      route: 'clients'
    },
    {
      icon: 'account_circle',
      title: 'Usuarios',
      route: 'users'
    },
    {
      icon: 'delete',
      title: 'Papelera',
      route: 'trashed?m=orders'
    }
    /* ,{
      icon: 'mail_outline',
      title: 'Plantillas de Emails',
      route: 'mails'
    } */
  ]
}

export default {
  namespaced: true,
  state: {sidemenu: app.sidemenu},
  mutations: { },
  actions: { },
  getters: { }
}
