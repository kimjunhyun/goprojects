using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Threading;

//Using namespaces 
using System.Data;
using MySql.Data.MySqlClient;
using System.Configuration;

namespace RabbitTest
{
    public class Receive
    {
        MainWindow m_pMainWnd;
        Thread recvThread;
        string QueueName = "안녕하세요";
        public string msg1;

        MySqlConnection conn = new MySqlConnection(ConfigurationManager.ConnectionStrings["ConnectionString"].ConnectionString);      

        public Receive(MainWindow pWnd)
        {
            m_pMainWnd = pWnd;
            recvThread = new Thread(new ThreadStart(RecvThread));
            recvThread.Start();
        }
        void RecvThread()
        {
            try
            {
                conn.Open();
            }
            catch (MySqlException ex)
            {
            }
            var factory = new ConnectionFactory();
            factory.HostName = "localhost";
            factory.UserName = "guest";
            factory.Password = "guest";
            using (var connection = factory.CreateConnection())
            using (var channel = connection.CreateModel())
            {
                // When reading from a persistent queue, you need to tell that to your consumer
                const bool durable = true;
                channel.QueueDeclare(QueueName, durable, false, false, null);

                var consumer = new QueueingBasicConsumer(channel);

                // turn auto acknowledge off so we can do it manually. This is so we don't remove items from the queue until we're perfectly happy
                const bool autoAck = false;
                channel.BasicConsume(QueueName, autoAck, consumer);

                string query2;
                while (true)
                {
                    var ea = (BasicDeliverEventArgs)consumer.Queue.Dequeue();

                    byte[] body = ea.Body;
                    string message = System.Text.Encoding.UTF8.GetString(body);
                    System.Console.WriteLine(" [x] Processing {0}", message);

                    // Acknowledge message received and processed
                    System.Console.WriteLine(" Processed ", message);

                    channel.BasicAck(ea.DeliveryTag, false);

                    msg1 = message;
/*
                    if (message.Length > 10)
                    {
                        query2 = message;
                        MySqlCommand cmdSpkSel4 = conn.CreateCommand();
                        cmdSpkSel4.CommandText = query2;
                        cmdSpkSel4.ExecuteNonQuery();
                    }
 * */
                }
                conn.Close();
            }
        }
        public void close()
        {
            recvThread.Abort();
        }
    }
}
