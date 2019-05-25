using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

using System.Windows.Threading;

using RabbitMQ.Client;

//Using namespaces 
using System.Data;
using MySql.Data.MySqlClient;
using System.Configuration;
namespace RabbitTest
{
    /// <summary>
    /// MainWindow.xaml에 대한 상호 작용 논리
    /// </summary>
    public partial class MainWindow : Window
    {
        Receive receive;

        int im_cnt = 0;

        public MainWindow()
        {
            InitializeComponent();
            receive = new Receive(this);
            System.Threading.Timer timer = new System.Threading.Timer(CallBack);
            Send send = new Send("localhost", "hello",  "test");
            timer.Change(0, 1000);
        }

        delegate void TimerEventFiredDelegate();
        void CallBack(object state)
        {
            Dispatcher.BeginInvoke(new TimerEventFiredDelegate(Work));
        }

        private void Work()
        {
        //수행해야할 작업(UI Thread 핸들링 가능)
            im_cnt++;
            sendTxt.Text = im_cnt.ToString();
            receiveTxt.Text = receive.msg1;
            Send send = new Send("localhost", "hello", sendTxt.Text+"ddd");
        }


        private void sendBtn_Click(object sender, RoutedEventArgs e)
        {
            Send send = new Send("localhost", "hello", "received");
        }


        private void Window_Closing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            receive.close();
        }
    }
}
