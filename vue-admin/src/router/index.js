import Vue from 'vue'
import Router from 'vue-router'

/**
 * Components
 */

import login from '@/components/login'
import home from '@/components/home'
import orders from '@/components/orders'
import activities from '@/components/activities'
import briefs from '@/components/briefs'
import clients from '@/components/clients'
import countries from '@/components/countries'
import coupons from '@/components/coupons'
import currencies from '@/components/currencies'
import gateways from '@/components/gateways'
import locations from '@/components/locations'
import mails from '@/components/mails'
import paymentsMethods from '@/components/payments_methods'
import sectors from '@/components/sectors'
import services from '@/components/services'
import users from '@/components/users'
import portfolios from '@/components/portfolios'
import trashed from '@/components/trashed'

Vue.use(Router)

var instance = new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: login
    },
    {
      path: '/home',
      name: 'home',
      component: home
    },
    {
      path: '/orders',
      name: 'orders',
      component: orders,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/activities',
      name: 'activities',
      component: activities,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/briefs',
      name: 'briefs',
      component: briefs,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/clients',
      name: 'clients',
      component: clients,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/countries',
      name: 'countries',
      component: countries,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/coupons',
      name: 'coupons',
      component: coupons,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/currencies',
      name: 'currencies',
      component: currencies,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/gateways',
      name: 'gateways',
      component: gateways,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/locations',
      name: 'locations',
      component: locations,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/mails',
      name: 'mails',
      component: mails,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/payments-methods',
      name: 'payments-methods',
      component: paymentsMethods,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/sectors',
      name: 'sectors',
      component: sectors,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/services',
      name: 'services',
      component: services,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/users',
      name: 'users',
      component: users,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/portfolios',
      name: 'portfolios',
      component: portfolios,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/trashed',
      name: 'trashed',
      component: trashed,
      props: (route) => ({ model: route.query.m, search: route.query.q })
    }
  ]
})

export default instance
