import Layout from '@/layout/index'

const myGameRouter = {
  path: '/go-plot',
  component: Layout,
  redirect: '/go-plot/index',
  name: 'go-plot',
  meta: {
    title: 'goPlot',
    icon: 'star'
  },
  children: [
    {
      path: 'index',
      component: () => import('@/views/goplot/index'),
      name: 'index',
      meta: { title: 'index', icon: 'user', noCache: true }
    },
    {
      path: 'plotPc',
      name: 'PlotPc',
      component: () => import('@/views/goplot/plotPc'),
      meta: {
        title: '批图节点列表',
        icon: 'tree',
        roles: ['admin']
      }
    },
    {
      path: 'plotUsrPc',
      name: 'PlotUserPc',
      component: () => import('@/views/goplot/plotUserPc'),
      meta: {
        title: '批图节点',
        icon: 'tree'
      }
    },
    {
      path: 'plotDisk',
      name: 'PlotJob',
      component: () => import('@/views/goplot/plotDisk'),
      meta: { title: '批图磁盘', icon: 'tree' }
    },
    {
      path: 'upLoadFile',
      name: 'UpLoadFile',
      component: () => import('@/views/goplot/upLoadFile'),
      meta: {
        title: '上传文件',
        icon: 'tree',
        roles: ['admin']
      }
    },
    {
      path: 'global_list',
      name: 'globalList',
      component: () => import('@/views/goplot/globalList'),
      meta: {
        title: '全局变量',
        icon: 'tree',
        roles: ['admin']
      }
    }
  ]
}

export default myGameRouter
