import Layout from '@/layout/index'

const jdSecKillRouter = {
  path: '/jd-sec-kill',
  component: Layout,
  redirect: '/jd-sec-kill/index',
  name: 'jd-sec-kill',
  meta: {
    title: 'jd-seckill',
    icon: 'star'
  },
  children: [
    {
      path: 'index',
      component: () => import('@/views/jdseckill/index'),
      name: 'index',
      meta: { title: 'index', icon: 'user', noCache: true ,roles: ['admin','jdkill']}
    },
    {
      path: 'miaoshaList',
      name: 'miaoshaList',
      component: () => import('@/views/jdseckill/miaoshaList/index'),
      meta: {
        title: '秒杀列表',
        icon: 'tree',
        roles: ['admin','jdkill']
      }
    },
    {
      path: 'itemInfo',
      name: 'ItemInfo',
      component: () => import('@/views/jdseckill/miaoshaItem/index'),
      meta: {
        title: '秒杀物品信息',
        icon: 'tree',
        roles: ['admin','jdkill']
      }
    },
    {
      path: 'panicBuying',
      name: 'ItemInfo',
      component: () => import('@/views/jdseckill/panicBuyingList/index'),
      meta: {
        title: '抢购设置',
        icon: 'tree',
        roles: ['admin','jdkill']
      }
    }
  ]
}

export default jdSecKillRouter
